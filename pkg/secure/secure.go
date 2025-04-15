package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

type Secure interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherHex string) (string, error)
}

type secure struct {
	gcm cipher.AEAD
}

func New(encryptionKey string) (sec Secure, err error) {
	key, err := hex.DecodeString(encryptionKey)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	sec = &secure{
		gcm: gcm,
	}
	return
}

func (s *secure) Encrypt(plainText string) (enc string, err error) {
	nonce := make([]byte, s.gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return
	}

	cipherText := s.gcm.Seal(nonce, nonce, []byte(plainText), nil)
	enc = hex.EncodeToString(cipherText)
	return
}

func (s *secure) Decrypt(cipherHex string) (dec string, err error) {
	decodedCipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return
	}

	decryptedData, err := s.gcm.Open(nil, decodedCipherText[:s.gcm.NonceSize()], decodedCipherText[s.gcm.NonceSize():], nil)
	if err != nil {
		return
	}

	dec = string(decryptedData)
	return
}
