package labutils

import (
	"regexp"
	"strings"
	"unsafe"
)

// TrimHtml 过滤html代码
func TrimHtml(src string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// StringToBytes
// converts string to byte slice without a memory allocation.
// 将字符串转换为字节切片
func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return []byte("")
	}
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString
// converts byte slice to string without a memory allocation.
// 将字节切片转换为字符串
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// BytesSplit
// split the byte slice by the specified size.
// 按照指定长度分割字节切片
func BytesSplit(buf []byte, size int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/size+1)
	for len(buf) >= size {
		chunk, buf = buf[:size], buf[size:]
		chunks = append(chunks, chunk)
	}
	return chunks
}
