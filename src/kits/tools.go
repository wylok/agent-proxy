package kits

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"os"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Max(value []uint64) uint64 {
	var max uint64
	for _, val := range value {
		if val > max {
			max = val
		}
	}
	return max
}
func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
func CheckFile(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
