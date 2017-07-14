package spafs_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	spafs "github.com/koron/go-spafs"
)

func TestFileServer(t *testing.T) {
	ts := httptest.NewServer(spafs.FileServer(http.Dir("testdata")))
	defer ts.Close()

	// check functions:
	_get := func(path string) *http.Response {
		r, err := http.Get(ts.URL + path)
		if err != nil {
			t.Fatalf("GET %q failed: %s", path, err)
		}
		return r
	}
	_check_sc := func(path string, code int) *http.Response {
		r := _get(path)
		if r.StatusCode != code {
			t.Fatalf("unexpected status code for GET %q: expected=%d actual=%d", path, code, r.StatusCode)
		}
		return r
	}
	_check := func(path string, fn func(act string)) {
		r := _check_sc(path, 200)
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read body of %q: %s", path, err)
		}
		fn(string(b))
	}
	eq := func(path string, content string) {
		_check(path, func(act string) {
			if act != content {
				t.Fatalf("content of %q is not matched: expected=%q actual=%q", path, content, act)
			}
		})
	}
	not := func(path string, contents []string) {
		_check(path, func(act string) {
			for _, bad := range contents {
				if act == bad {
					t.Fatalf("content of %q shouldn't matched: %q", path, bad)
				}
			}
		})
	}
	notfound := func(path string) {
		r := _check_sc(path, 404)
		defer r.Body.Close()
	}

	known := []string{
		"00000001\n",
		"00000002\n",
		"00000003\n",
		"00000004\n",
		"00000005\n",
		"00000006\n",
	}

	not("/", known)
	not("/index.html", known)
	eq("/test.js", "00000001\n")

	eq("/foo/", "00000002\n")
	eq("/foo", "00000002\n")
	eq("/foo/foo.js", "00000003\n")
	eq("/bar/bar.js", "00000004\n")

	notfound("/not")
	notfound("/not/")
	notfound("/not/found")
	notfound("/not/found/")

	eq("/foo/bar/", "00000002\n")
	eq("/foo/bar", "00000002\n")
	eq("/foo/qux/", "00000002\n")
	eq("/foo/qux/baz", "00000002\n")
	eq("/foo/baz/", "00000005\n")
	eq("/foo/baz", "00000005\n")
	eq("/foo/baz/baz.js", "00000006\n")
	eq("/foo/baz/qux/", "00000005\n")

	notfound("/bar/foo/")
	notfound("/bar/baz/")
	not("/bar/", known)
	not("/bar", known)
}
