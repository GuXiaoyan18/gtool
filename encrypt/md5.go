package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetMd5Hash(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenCipherPasswd(password, salt string) string {
	h := md5.New()
	io.WriteString(h, password)
	io.WriteString(h, fmt.Sprintf("%x", h.Sum(nil)) + salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
