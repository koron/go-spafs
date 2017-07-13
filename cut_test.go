package spafs

import "testing"

func TestCut(t *testing.T) {
	ok := func(s string, expected string) {
		v := cut(s)
		if v != expected {
			t.Fatalf("cut failed: %q expected=%q actual=%q", s, expected, v)
		}
	}
	ok("/foo/bar/", "/foo/")
	ok("/foo/bar", "/foo/")
	ok("/foo/", "/")
	ok("/foo", "/")
}
