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
	get := func(path string) (*http.Response, error) {
		return http.Get(ts.URL + path)
	}
	ok := func(path string, content string) {
		r, err := get(path)
		if err != nil {
			t.Fatalf("get %q failed: %s", path, err)
		}
		defer r.Body.Close()
		if r.StatusCode != 200 {
			t.Fatalf("get %q ends with %d (not 200)", path, r.StatusCode)
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read body of %q: %s", path, err)
		}
		s := string(b)
		if s != content {
			t.Fatalf("content of %q is not matched: expected=%q actual=%q", path, content, s)
		}
	}

	ok("/index.html", "00000000\n")
	ok("/", "00000000\n")
	ok("/not/found", "00000000\n")
	//ok("/not/found/", "00000000\n")
	ok("/test.js", "00000001\n")
}
