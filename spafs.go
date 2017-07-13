package spafs

import (
	"net/http"
	"os"
	"path"
	"strings"
)

type fileHandler struct {
	root http.FileSystem
}

// FileServer returns a handler that serves HTTP requests with the contents of
// the file system rooted at root, and optimized for single page application
// (SPA).
func FileServer(root http.FileSystem) http.Handler {
	return &fileHandler{root: root}
}

func (fh *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	for upath != "/" {
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

	http.ServeFile(w, r, upath)
}

func cut(name string) string {
	dir, _ := path.Split(name)
	return dir
}
