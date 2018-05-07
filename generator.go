package goinsta

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strings"
)

const (
	volatileSeed = "12345"
)

func generateMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateHMAC(text, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateDeviceID(seed string) string {
	hash := generateMD5Hash(seed + volatileSeed)
	return "android-" + hash[:16]
}

func generateUUID(replace bool) string {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	tUUID := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	if replace {
		return strings.Replace(tUUID, "-", "", -1)
	}
	return tUUID
}

func generateSignature(data string) map[string]string {
	return map[string]string{
		"ig_sig_key_version": fmt.Sprintf(
			"%s&signed_body=%s.%s",
			goInstaSigKeyVersion,
			generateHMAC(data, goInstaIGSigKey),
			url.QueryEscape(data),
		),
	}
}
