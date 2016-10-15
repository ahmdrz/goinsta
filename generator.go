package goinsta

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"

	"github.com/satori/go.uuid"
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
	volatile_seed := "12345"
	hash := generateMD5Hash(seed + volatile_seed)
	return "android-" + hash[:16]
}

func generateUUID(replace bool) string {
	u1 := uuid.NewV4()
	if replace {
		return strings.Replace(u1.String(), "-", "", -1)
	}
	return u1.String()
}

func generateSignature(data string) string {
	return "ig_sig_key_version=" + GOINSTA_SIG_KEY_VERSION + "&signed_body=" + generateHMAC(data, GOINSTA_IG_SIG_KEY) + "." + url.QueryEscape(data)
}
