package server

import (
	"encoding/json"
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
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req encryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.store.Add(req.Key, req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) DecryptText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req decryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plainText, err := h.store.Remove(req.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := decryptResponse{PlainText: plainText}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

type encryptRequest struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}

type decryptRequest struct {
	Key string `json:"key"`
}

type decryptResponse struct {
	PlainText string `json:"plainText"`
}
