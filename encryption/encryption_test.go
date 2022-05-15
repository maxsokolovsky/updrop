package encryption_test

import (
	"testing"

	"github.com/maxsokolovsky/updrop/encryption"
)

func TestEncryptDecrypt(t *testing.T) {
	t.Parallel()

	const key = "helloworld1234ai"
	encrypter, err := encryption.NewEncrypter(key)
	if err != nil {
		t.Fatal(err.Error())
	}

	const plainText = "hello there"
	cipherText, err := encrypter.Encrypt(plainText)
	if err != nil {
		t.Fatalf("failed to encrypt: %v", err)
	}

	decrypted, err := encrypter.Decrypt(cipherText)
	if err != nil {
		t.Fatalf("failed to decrypt: %v", err)
	}

	if plainText != decrypted {
		t.Fatal("original plaintext and decrypted value are not equal")
	}
}
