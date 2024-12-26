package main

import (
	"fmt"
	"testing"
)

// 12246290580013200379
func TestEncodeAmountm(t *testing.T) {
	M := "0000100101111101101001001111101101110101"
	amountm := EncodeAmountm(M)
	fmt.Println(amountm)
}
