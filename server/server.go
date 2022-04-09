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

	r.HandleFunc("/encrypt", h.EncryptText).Methods("POST")
	r.HandleFunc("/decrypt", h.DecryptText).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))

	return &http.Server{
		Addr:    c.Addr,
		Handler: r,
	}
}
