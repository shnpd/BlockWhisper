package share

import (
	"fmt"
	"log"
	"strconv"
)

func XorBinaryString(a, b string) string {
	length := len(a)
	ret := ""
	for i := 0; i < length; i++ {
		if a[i] == b[i] {
			ret += "0"
		} else {
			ret += "1"
		}
	}
	return ret
}
func Byte2binary(b []byte) string {
	binaryStr := ""
	for _, v := range b {
		binaryStr += fmt.Sprintf("%08b", v)
	}
	return binaryStr
}

func Binary2byte(b string) []byte {
	var ret []byte
	for i := 0; i < len(b); i += 8 {
		value, err := strconv.ParseUint(b[i:i+8], 2, 8)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, byte(value))
	}
	return ret
}
