package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 這個函數 EncodeMD5 是用來計算輸入字串的 MD5 雜湊（hash）值，並將這個雜湊值轉換成十六進制字串表示
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
