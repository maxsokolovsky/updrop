package server

import (
	"errors"

	"github.com/maxsokolovsky/updrop/encryption"
)

type Store interface {
	Add(string, string) (string, error)
	Remove(string) (string, error)
	Size() int
}

type singleValueStore struct {
	value string
}

func NewSingleValueStore() Store {
	return &singleValueStore{}
}

func (s *singleValueStore) Add(key, value string) (string, error) {
	var cipherText string
	var err error

	cipherText, err = encryption.Encrypt(key, value)
	if err != nil {
		return "", err
	}
	s.value = cipherText
	return s.value, nil
}

func (s *singleValueStore) Remove(key string) (string, error) {
	if s.Size() == 0 {
		return "", errors.New("store is empty")
	}

	var plainText string
	var err error

	plainText, err = encryption.Decrypt(key, s.value)
	if err != nil {
		return "", err
	}
	s.value = ""
	return plainText, nil
}

func (s *singleValueStore) Size() int {
	return len(s.value)
}
