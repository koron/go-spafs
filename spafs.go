package spafs

import (
	"net/http"
	"os"
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

	for upath != "/" && upath != "" {
		f, err := fh.root.Open(upath)
		if err != nil {
			if os.IsNotExist(err) {
				upath = cut(upath)
				continue
			}
			break
		}
		f.Close()
		break
	}

	r.URL.Path = upath
	fh.fs.ServeHTTP(w, r)
}
