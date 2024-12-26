package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 示例 M 序列
	M := []int{1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0}

	// 编码为 Amountm
	amountm := encodeAmountm(M)

	fmt.Println("M:", M)
	fmt.Println("Amountm:", amountm)

	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
	M = decodeAmountm(amountm, 40)
	fmt.Println("Amountm:", amountm)
	fmt.Println("M:", M)

}
