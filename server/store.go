package server

import (
	"errors"

	"github.com/maxsokolovsky/updrop/encryption"
)

type Store interface {
	Add(string, string) (string, error)
	Remove(string, string) (string, error)
	Size() int
}

type singleValueStore struct {
	value         string
	encrypter     encryption.Encrypter
	serverWideKey bool
}

func NewSingleValueStore(key string, serverWideKey bool) *singleValueStore {
	var encrypter encryption.Encrypter
	var err error

	if serverWideKey {
		encrypter, err = encryption.NewEncrypter(key)
		if err != nil {
			panic(err)
		}
	}
	return &singleValueStore{
		encrypter:     encrypter,
		serverWideKey: serverWideKey,
	}
}

func (s *singleValueStore) Add(key, value string) (string, error) {
	var cipherText string
	var err error

	if s.serverWideKey {
		cipherText, err = s.encrypter.Encrypt(value)
	} else {
		cipherText, err = encryption.Encrypt(key, value)
	}
	if err != nil {
		return "", err
	}
	s.value = cipherText
	return s.value, nil
}

func (s *singleValueStore) Remove(key, value string) (string, error) {
	if s.Size() == 0 {
		return "", errors.New("store is empty")
	}

	var plainText string
	var err error

	if s.serverWideKey {
		plainText, err = s.encrypter.Decrypt(s.value)
	} else {
		plainText, err = encryption.Decrypt(key, value)
	}

	if err != nil {
		return "", err
	}
	s.value = ""
	return plainText, nil
}

func (s *singleValueStore) Size() int {
	return len(s.value)
}
