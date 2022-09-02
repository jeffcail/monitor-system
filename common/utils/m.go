package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5s
func Md5s(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	return md5Str
}
