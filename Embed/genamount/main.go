package main

import (
	mycrypto "blockwhisper/crypto"
	"blockwhisper/share"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "normal_amount.xlsx" // 替换为你的 Excel 文件路径
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer func() {
		// 确保文件在程序结束后关闭
		if err := f.Close(); err != nil {
			log.Fatalf("无法关闭文件: %v", err)
		}
	}()
	sheetName := "Sheet1"
	col, err := f.GetCols(sheetName)
	lengths := col[1][:100000]

	// 1. 选择输入输出地址
	//inputAddr := share.AddressPool[rand.IntN(len(share.AddressPool))]
	inputAddr := share.AddressPool[0]
	//outputAddr := share.AddressPool[rand.IntN(len(share.AddressPool))]
	//	2.
	for t := 0; t < 200; t++ {
		msg := "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
		cipher, err := mycrypto.Encrypt([]byte(msg), share.Key)
		if err != nil {
			log.Fatal("encrypto error:", err)
		}
		binaryCipher := share.Byte2binary(cipher)
		p := 32
		var Ms []string
		var M string
		M2 := share.Byte2binary([]byte(inputAddr))[:p]

		for i := 0; i < len(binaryCipher); i += p {
			var M1 string
			if i+p >= len(binaryCipher) {
				M1 = binaryCipher[i:]
			} else {
				M1 = binaryCipher[i : i+p]
			}
			M = share.XorBinaryString(M1, M2)
			Ms = append(Ms, M)
		}
		// 3 金额编码
		var amounts []string
		for _, v := range Ms {
			amount := EncodeAmountm(v)

			lenAmount := lengths[rand.IntN(len(lengths))]
			lenAmountInt, _ := strconv.Atoi(lenAmount)
			lenAmountInt = minInt(lenAmountInt, len(amount))

			amount = amount[:lenAmountInt]
			amounts = append(amounts, amount)
		}
		//	保存金额
		saveAmount(amounts, "covertAmount.xlsx")
	}

}
func saveAmount(a []string, filepath string) {
	_, err := os.Stat(filepath)
	if err != nil {
		f := excelize.NewFile()
		// 添加一个默认的工作表
		f.NewSheet("Sheet1")
		// 保存文件
		if err := f.SaveAs(filepath); err != nil {
			log.Fatal(err)
		}
	}
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	sheetName := "Sheet1"
	cols, _ := f.GetCols(sheetName)
	column := "A" // 要写入的列，例如 A 列
	var start int
	if len(cols) == 0 {
		start = 0
	} else {
		start = len(cols[0])
	}
	for i, value := range a {
		cell := fmt.Sprintf("%s%d", column, start+i+1) // 拼接单元格名称，例如 A1, A2, ...
		if err := f.SetCellValue(sheetName, cell, value); err != nil {
			log.Fatalf("写入单元格 %s 失败: %v", cell, err)
		}
	}
	//f.SetActiveSheet(index)

	// 保存 Excel 文件
	if err := f.SaveAs(filepath); err != nil {
		log.Fatalf("保存文件失败: %v", err)
	}

	fmt.Printf("数据已成功保存到 %s 的 %s 列\n", filepath, column)
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
