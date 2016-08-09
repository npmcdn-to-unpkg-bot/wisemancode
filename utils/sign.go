package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

//微信公众号 验签
//原理：首先生成一个切片
//然后排序
//最后组装成字符串
//转化成字节切片，使用sha1生成一个串
func Sign(token, timestamp, nonce string) (signature string) {
	strs := []string{token, timestamp, nonce}
	var buf []byte = StringUnionSort(strs)
	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

func MsgSign(token, timestamp, nonce, enc string) (signature string) {
	strs := []string{token, timestamp, nonce, enc}
	var buf []byte = StringUnionSort(strs)
	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
func StringUnionSort(s []string) []byte {
	if len(s) == 0 {
		return nil
	}
	sort.Strings(s)
	var str string = strings.Join(s[:], "")
	buf := make([]byte, len(str))
	buf = append(buf, str...)
	return buf
}
