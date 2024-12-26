package main

import "strconv"

// DecodeAmountm 根据编码 Amountm 解码为位序列 M
func DecodeAmountm(amountm string, length int) string {
	M := make([]int, length) // 初始化位序列，全为 1
	for i := 0; i < length; i++ {
		M[i] = 1
	}
	pos := 0 // 当前处理的起始位置（基于 10 位分段调整）

	// 逐字符解析 Amountm
	for i := 0; i < len(amountm); {
		ch := string(amountm[i])

		if ch == "0" && amountm[i+1] == '0' {
			pos += 10
			M[pos-1] = 0
			i += 2
			continue
		}
		if ch == "0" {
			// 遇到分隔符，跨越十的倍数
			pos += 10
			i++
			continue
		}

		// 当前数字
		num1, _ := strconv.Atoi(ch)
		// 后一个数字
		var num2 int
		// 检查是否有第二个数字,没有则解码单个零
		if i+1 < len(amountm) && amountm[i+1] != '0' {
			num2, _ = strconv.Atoi(string(amountm[i+1]))
		} else {
			// 解码单个零
			M[pos+num1-1] = 0
			i++
			continue
		}
		// 如果后一个数比前一个小，那么这两个数字一定表示连续零，如果后面的比前面的大，那么前面的一定表示单个零，后面的不一定，可能表示单个零也可能与其后面的数字一起表示连续零
		if num2 <= num1 {
			// 解码连续零
			start := pos + num1
			for j := 0; j < num2; j++ {
				M[start+j-1] = 0
			}
			i += 2
		} else {
			M[pos+num1-1] = 0
			i++
		}
	}
	ret := ""
	for _, v := range M {
		ret += strconv.Itoa(v)
	}
	return ret
}
