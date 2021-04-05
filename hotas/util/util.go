package util

import (
	"encoding/binary"
)

func GetAxis(b []byte) int32 {
	return int32(binary.LittleEndian.Uint16(b))
}

func GetBool(b byte, i int) bool {
	if i < 0 || i > 7 {
		return false
	}
	return b>>(8-i-1)&1 == 1
}
