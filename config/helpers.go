package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
	"time"
)

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func ParseDate(dateStr string) (time.Time, error) {
	layouts := []string{"2006-01-02", "02/01/2006", "January 2, 2006"}
	for _, layout := range layouts {
		parsedDate, err := time.Parse(layout, dateStr)
		if err == nil {
			return parsedDate, nil
		}
	}
	return time.Time{}, errors.New("invalid date format")
}

func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}

func ContainsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}