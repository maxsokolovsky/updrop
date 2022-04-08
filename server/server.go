package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/maxsokolovsky/updrop/config"
)

func New(c config.Config) *http.Server {
	s := NewSingleValueStore(c.Key, c.ServerWideKey)
	h := NewHandler(s)
	r := mux.NewRouter()

	r.HandleFunc("/new", h.EncryptText).Methods("POST")
	r.HandleFunc("/", h.DecryptText).Methods("POST")

	return &http.Server{
		Addr:    c.Addr,
		Handler: r,
	}
}
