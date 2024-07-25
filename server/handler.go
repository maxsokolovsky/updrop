package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	store Store
}

func NewHandler(s Store) *Handler {
	return &Handler{
		store: s,
	}
}

func (h *Handler) EncryptText(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) DecryptText(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, ErrEmptyStore) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	resp := decryptResponse{PlainText: plainText}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
