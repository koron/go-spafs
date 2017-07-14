package spafs

import (
	"net/http"
	"os"
	"path"
	"strings"
)

type fileHandler struct {
	root http.FileSystem
	fs   http.Handler
}

// FileServer returns a handler that serves HTTP requests with the contents of
// the file system rooted at root, and optimized for single page application
// (SPA).
func FileServer(root http.FileSystem) http.Handler {
	return &fileHandler{
		root: root,
		fs:   http.FileServer(root),
	}
}

func (fh *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
	}

	f, err := fh.root.Open(upath)
	if err == nil {
		f.Close()
	}

	// search index.html in ancestor directories and return it if exists.
	if err != nil && os.IsNotExist(err) {
		for upath != "/" {
			upath = cut(upath)
			f, err := fh.root.Open(path.Join(upath, "index.html"))
			switch {
			case err == nil:
				f.Close()
				fallthrough
			case !os.IsNotExist(err):
				r.URL.Path = upath
				fh.fs.ServeHTTP(w, r)
				return
			}
		}
	}

	fh.fs.ServeHTTP(w, r)
}
