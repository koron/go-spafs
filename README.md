# SPA Specialized File Server

spafs is a file server which specialized for SPA (single page application).
This serves content of index.html in nearest ancestral directory for
unavailable path.

## How to use

```go

import (
	"http"
	spafs "github.com/koron/go-spafs"
)

fs := spafs.FileServer(http.Dir("./assets"))
http.Handle("/", fs)
```
