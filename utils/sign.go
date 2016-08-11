package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"sort"
	"strings"
	"wisemancode/log"
)

//微信公众号 验签
//原理：首先生成一个切片
//然后排序
//最后组装成字符串
//转化成字节切片，使用sha1生成一个串
func Sign(token, timestamp, nonce string) (signature string, err error) {
	if len(token) == 0 {
		return "", errors.New("token is null")
	}
	if len(timestamp) == 0 {
		return "", errors.New("timestamp is null")
	}
	if len(nonce) == 0 {
		return "", errors.New("nonce is null")
	}
	strs := []string{token, timestamp, nonce}

	var buf []byte = StringUnionSort(strs)
	hashsum := sha1.Sum(buf)
	signature = hex.EncodeToString(hashsum[:])
	log.Logger.Info("计算签名串：" + signature)
	return signature, nil
}

func MsgSign(token, timestamp, nonce, enc string) (signature string) {
	strs := []string{token, timestamp, nonce, enc}
	var buf []byte = StringUnionSort(strs)
	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
func StringUnionSort(s []string) []byte {
	if len(s) == 0 {
		log.Logger.Error("验证签名字典排序错误，排序字符串nil")
		return nil
	}

	sort.Strings(s)
	log.Logger.Info("需要验证签名的数据：" + strings.Join(s, "|"))
	var str string = strings.Join(s[:], "")
	buf := make([]byte, len(str))
	buf = append(buf, str...)
	return buf
}
