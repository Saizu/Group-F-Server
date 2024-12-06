package main

import (
	"path/filepath"
	"net/http"
	"runtime"
)

func main() {
	_, path, _, _ := runtime.Caller(0)
	dir := filepath.Dir(path)

	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":8080", nil)
}
