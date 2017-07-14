# SPA Specialized File Server

spafs is a file server which specialized for SPA (single page application).
This serves content of index.html in nearest ancestral directory for
unavailable path.

## How to use

```go
package main

import (
    "http"
    spafs "github.com/koron/go-spafs"
)

func main() {
    fs := spafs.FileServer(http.Dir("./testdata"))
    err := http.ListenAndServe(":8080", fs)
    if err != nil {
        panic(err)
    }
}
```
