package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// GeneratePassword
func GeneratePassword(password, slat string) string {
	return Md5s(Md5s(password) + slat)
}

// ComparePassword
func ComparePassword(password, loginPass, slat string) bool {
	if password == "" || loginPass == "" {
		return false
	}

	n := GeneratePassword(loginPass, slat)

	if password != n {
		return false
	}
	return true
}

// Md5s
func Md5s(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	return md5Str
}
