package encrypter_test

import (
	"testing"

	"github.com/maxsokolovsky/updrop/encrypter"
)

func TestEncryptDecrypt(t *testing.T) {
	t.Parallel()
	const key = "helloworld1234ai"
	enc, err := encrypter.New(key)
	if err != nil {
		t.Fatal(err.Error())
	}

	const plainText = "hello there"
	cipherText, err := enc.Encrypt(plainText)
	if err != nil {
		t.Fatalf("failed to encrypt: %v", err)
	}

	decrypted, err := enc.Decrypt(cipherText)
	if err != nil {
		t.Fatalf("failed to decrypt: %v", err)
	}

	if plainText != decrypted {
		t.Fatal("original plaintext and decrypted value are not equal")
	}
}
