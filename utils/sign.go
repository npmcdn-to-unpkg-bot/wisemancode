package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"

	"io"
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

	return Sign1(token, timestamp, nonce)
}
func Sign1(params ...string) (signature string, err error) {
	if len(params) == 0 {
		log.Logger.Error("params参数为空")
		return "", errors.New("params参数为空")
	}
	log.Logger.Info("params参数进行字典排序")
	str := StringUnionSortToStr(params)
	log.Logger.Info("params参数进行字典排序结果 %s", str)
	log.Logger.Info("生成sha1摘要digest实例")
	s := sha1.New()
	log.Logger.Info("生成sha1摘要digest实例：%+v", s)
	log.Logger.Info("生成sha1摘要digest实例,写入字节数据")
	io.WriteString(s, str)
	log.Logger.Info("生成sha1摘要digest实例,写入字节数据 %+v", s)
	hashsum := s.Sum(nil)
	signature = hex.EncodeToString(hashsum[:])
	log.Logger.Info("计算签名串：" + signature)
	return signature, nil
}
func MsgSign(token, timestamp, nonce, enc string) (signature string, err error) {

	if len(token) == 0 {
		return "", errors.New("token is null")
	}
	if len(timestamp) == 0 {
		return "", errors.New("timestamp is null")
	}
	if len(nonce) == 0 {
		return "", errors.New("nonce is null")
	}

	if len(enc) == 0 {
		return "", errors.New("enc is null")
	}
	return Sign1(token, timestamp, nonce, enc)
}
func StringUnionSort(s []string) []byte {
	if len(s) == 0 {
		log.Logger.Error("验证签名字典排序错误，排序字符串nil")
		return nil
	}

	sort.Strings(s)
	log.Logger.Info("需要验证签名的数据：" + strings.Join(s, "|"))
	var str string = strings.Join(s[:], "")
	log.Logger.Info("需要签名str:" + str)
	buf := make([]byte, len(str))
	buf = append(buf, str...)
	return buf
}
func StringUnionSortToStr(s []string) string {
	if len(s) == 0 {
		log.Logger.Error("验证签名字典排序错误，排序字符串nil")
		return ""
	}

	sort.Strings(s)
	log.Logger.Info("需要验证签名的数据：" + strings.Join(s, "|"))
	var str string = strings.Join(s[:], "")
	return str
}
