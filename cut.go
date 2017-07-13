package spafs

import (
	"path"
	"strings"
)

func cut(name string) string {
	name = strings.TrimSuffix(name, "/")
	dir, _ := path.Split(name)
	return dir
}
