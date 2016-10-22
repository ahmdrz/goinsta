package goinsta

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"
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
	uuid, err := newUUID()
	if err != nil {
		return "cb479ee7-a50d-49e7-8b7b-60cc1a105e22" // default value when error occurred
	}
	if replace {
		return strings.Replace(uuid, "-", "", -1)
	}
	return uuid
}

func generateSignature(data string) string {
	return "ig_sig_key_version=" + GOINSTA_SIG_KEY_VERSION + "&signed_body=" + generateHMAC(data, GOINSTA_IG_SIG_KEY) + "." + url.QueryEscape(data)
}
