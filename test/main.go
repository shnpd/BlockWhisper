package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	t := time.Duration(26) * (100 * time.Millisecond)
	fmt.Println(t)
}
func byte2binary(b []byte) string {
	binaryStr := ""
	for _, v := range b {
		binaryStr += fmt.Sprintf("%08b", v)
	}
	return binaryStr
}

func binary2byte(b string) []byte {
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
