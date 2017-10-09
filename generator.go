package goinsta

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	"github.com/ahmdrz/goinsta/uuid"
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
	tempUUID, err := uuid.NewUUID()
	if err != nil {
		return "cb479ee7-a50d-49e7-8b7b-60cc1a105e22" // default value when error occurred
	}
	if replace {
		return strings.Replace(tempUUID, "-", "", -1)
	}
	return tempUUID
}

func generateSignature(data string) string {
	return fmt.Sprintf("ig_sig_key_version=%s&signed_body=%s.%s",
		GOINSTA_SIG_KEY_VERSION,
		generateHMAC(data, GOINSTA_IG_SIG_KEY),
		url.QueryEscape(data),
	)
}
