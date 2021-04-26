package goinsta

import (
	"strconv"
	"strings"
)

func MediaIDFromShortID(code string) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	media_id := 0
	for _, letter := range code {
		media_id = (media_id * 64) + strings.Index(alphabet, string(letter))
	}
	result := strconv.Itoa(media_id)

	return result
}
