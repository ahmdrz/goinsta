package store

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/ahmdrz/goinsta"
)

// Secret is main struct for strore functions
type Secret struct {
	c cipher.Block
}

func newSecret(key []byte) (*Secret, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &Secret{c}, nil
}

func (s *Secret) encryptAES(plaintext []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(s.c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (s *Secret) decryptAES(ciphertext []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(s.c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// Import goinsta.Instagram
func Import(input, key []byte) (*goinsta.Instagram, error) {
	_bytes, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return nil, err
	}
	secret, err := newSecret(key)
	if err != nil {
		return nil, err
	}
	_bytes, err = secret.decryptAES(_bytes)
	if err != nil {
		return nil, err
	}
	backupType := goinsta.BackupType{}
	writer := bytes.NewBuffer(_bytes)
	decoder := gob.NewDecoder(writer)
	decoder.Decode(&backupType)

	_cookiejar, _ := cookiejar.New(nil)
	u, _ := url.Parse(goinsta.GOINSTA_API_URL)

	tmp := make([]*http.Cookie, 0)
	for i := range backupType.Cookies {
		tmp = append(tmp, &backupType.Cookies[i])
	}
	_cookiejar.SetCookies(u, tmp)

	insta := &goinsta.Instagram{}
	insta.InstaType = backupType.InstaType
	insta.Cookiejar = _cookiejar

	return insta, nil
}

// Export goinsta.Instgram
func Export(insta *goinsta.Instagram, key []byte) ([]byte, error) {
	backupType := goinsta.BackupType{}
	backupType.InstaType = insta.InstaType
	backupType.Cookies = make([]http.Cookie, 0)

	u, _ := url.Parse(goinsta.GOINSTA_API_URL)
	for _, value := range insta.Cookiejar.Cookies(u) {
		backupType.Cookies = append(backupType.Cookies, *value)
	}

	writer := bytes.NewBufferString("")
	encoder := gob.NewEncoder(writer)
	encoder.Encode(backupType)

	secret, err := newSecret(key)
	if err != nil {
		return nil, err
	}
	_bytes, err := secret.encryptAES(writer.Bytes())
	if err != nil {
		return nil, err
	}
	result := base64.StdEncoding.EncodeToString(_bytes)
	return []byte(result), nil
}
