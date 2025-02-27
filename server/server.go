package server

import (
	"crypto/tls"
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/maxsokolovsky/updrop/config"
)

//go:embed static
var staticFiles embed.FS

func New(c config.Config) *http.Server {
	staticFS := fs.FS(staticFiles)
	htmlContent, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}

	s := NewSingleValueStore()
	h := NewHandler(s)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(htmlContent)))
	mux.HandleFunc("/encrypt", h.EncryptText)
	mux.HandleFunc("/decrypt", h.DecryptText)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}
	return &http.Server{
		Addr:           c.Addr,
		Handler:        mux,
		TLSConfig:      tlsConfig,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
}
