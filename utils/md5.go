package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

//MD5 产生md5随机数
func MD5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

//Rand 使用纳秒 和MD5产生一个随机串
func Rand() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rnd := rand.Int63()
	return MD5(strconv.FormatInt(nano, 10)) + MD5(strconv.FormatInt(rnd, 10))
}
