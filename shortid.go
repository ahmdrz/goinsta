package goinsta

import (
	"fmt"
	"strconv"
	"strings"
)

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

const base64UrlCharmap = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func MediaIDFromShortID(code string) (string, error) {
	strID := ""
	for i := 0; i < len(code); i++ {
		base64 := strings.Index(base64UrlCharmap, string(code[i]))
		str2bin := strconv.FormatInt(int64(base64), 2)
		sixbits := leftPad2Len(str2bin, "0", 6)
		strID = strID + sixbits
	}
	result, err := strconv.ParseInt(strID, 2, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", result), nil
}
