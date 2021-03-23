package goinsta

import "fmt"

const shortIDLen = 11

func MediaIDFromPrivateID(code string) (string, error) {
	if len(code) < shortIDLen {
		return "", fmt.Errorf("private id is shorter than %v characters", shortIDLen)
	}
	shortID := code[:11]
	return MediaIDFromShortID(shortID)
}
