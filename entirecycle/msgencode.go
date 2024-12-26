package main

import (
	"strconv"
	"strings"
)

// encodeAmountm 根据规则将 M 编码为 Amountm
func EncodeAmountm(M string) string {
	var amountm []string
	zeroStart := -1
	zeroCount := 0

	for i := 0; i < len(M); i++ {
		bit := M[i] - '0'
		if bit == 0 {
			// 本次出现的第一个零
			if zeroStart == -1 {
				zeroStart = (i + 1) % 10 // 位序从 1 开始
				zeroCount = 1
			} else {
				// 连续的零
				zeroCount++
				// 后一位大于前一位先将之前的编码
				if zeroCount > zeroStart {
					if zeroCount == 2 {
						amountm = append(amountm, strconv.Itoa(zeroStart))
					} else {
						amountm = append(amountm, strconv.Itoa(zeroStart), strconv.Itoa(zeroCount-1))
					}
					i--
					zeroStart = -1
					zeroCount = 0
				}
			}
		} else {
			// 说明在此之前是一段零，要添加编码
			if zeroStart != -1 {
				// 添加零段编码
				if zeroCount == 1 {
					amountm = append(amountm, strconv.Itoa(zeroStart))
				} else {
					amountm = append(amountm, strconv.Itoa(zeroStart), strconv.Itoa(zeroCount))
				}
				zeroStart = -1
				zeroCount = 0
			}
		}

		// 检查是否需要插入“0”
		if (i+1)%10 == 0 {
			// 如果需要插入则先将之前的值插入金额
			if zeroStart != -1 {
				if zeroCount == 1 {
					amountm = append(amountm, strconv.Itoa(zeroStart))
				} else {
					amountm = append(amountm, strconv.Itoa(zeroStart), strconv.Itoa(zeroCount))
				}
				zeroStart = -1
				zeroCount = 0
			}
			amountm = append(amountm, "0")
		}
	}

	// 检查最后一段零
	if zeroStart != -1 {
		if zeroCount == 1 {
			amountm = append(amountm, strconv.Itoa(zeroStart))
		} else {
			amountm = append(amountm, strconv.Itoa(zeroStart), strconv.Itoa(zeroCount))
		}
	}

	// 移除最后可能多余的“0”
	if len(amountm) > 0 && amountm[len(amountm)-1] == "0" {
		amountm = amountm[:len(amountm)-1]
	}
	return strings.Join(amountm, "")
}
