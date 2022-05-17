package labutils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// Md5 获取字符串的md5值
// @return 小写32位
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Md5File 获取文件的md5值
func Md5File(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
