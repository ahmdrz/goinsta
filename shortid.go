package goinsta

import (
	"fmt"
	"strconv"
	"strings"
)

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func bin2int(binStr string) string {
	result, _ := strconv.ParseInt(binStr, 2, 64)
	return strconv.FormatInt(result, 10)
}

// Base64UrlCharmap - all posible characters
const Base64UrlCharmap = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func FromCode(code string) string {

	base2 := ""
	for i := 0; i < len(code); i++ {
		base64 := strings.Index(Base64UrlCharmap, string(code[i]))
		str2bin := strconv.FormatInt(int64(base64), 2)
		sixbits := leftPad2Len(str2bin, "0", 6)
		base2 = base2 + sixbits
	}

	return bin2int(base2)
}
