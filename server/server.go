package server

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/maxsokolovsky/updrop/config"
)

//go:embed static
var staticFiles embed.FS

func New(c config.Config) *http.Server {
	s := NewSingleValueStore()
	h := NewHandler(s)

	var staticFS = fs.FS(staticFiles)
	htmlContent, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(htmlContent)))
	mux.HandleFunc("/encrypt", h.EncryptText)
	mux.HandleFunc("/decrypt", h.DecryptText)

	return &http.Server{
		Addr:    c.Addr,
		Handler: mux,
	}
}
