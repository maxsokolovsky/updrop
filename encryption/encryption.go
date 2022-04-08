package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var ErrInvalidKeyLength = errors.New("invalid key length - must be either 16, 24, or 32 bytes")

type Encrypter interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type encrypter struct {
	key string
}

func NewEncrypter(key string) (*encrypter, error) {
	switch len(key) {
	case 8, 16, 24:
	default:
		return nil, ErrInvalidKeyLength
	}
	return &encrypter{key}, nil
}

func (e *encrypter) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(e.key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// Encode to base64 to make cipher-text look prettier than literal binary data.
	cipherText := base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(text), nil))
	return string(cipherText), nil
}

func (e *encrypter) Decrypt(cipherText string) (string, error) {
	block, err := aes.NewCipher([]byte(e.key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	cipherText = string(decoded)
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func Encrypt(key, value string) (string, error) {
	enc, err := NewEncrypter(key)
	if err != nil {
		return "", err
	}
	return enc.Encrypt(value)
}

func Decrypt(key, value string) (string, error) {
	enc, err := NewEncrypter(key)
	if err != nil {
		return "", err
	}
	return enc.Decrypt(value)
}