package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct {
	store Store
}

func NewHandler(s Store) *handler {
	return &handler{
		store: s,
	}
}

func (h *handler) EncryptText(w http.ResponseWriter, r *http.Request) {
	var req encryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cipherText, err := h.store.Add(req.Key, req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s\n", cipherText)
}

func (h *handler) DecryptText(w http.ResponseWriter, r *http.Request) {
	var req decryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plainText, err := h.store.Remove(req.Key, req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", plainText)
}

type encryptRequest struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}

type decryptRequest struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}
