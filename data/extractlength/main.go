package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	var lencount = make(map[int]int)
	// 打开 Excel 文件
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
	for i := 0; i < 100000; i++ {
		lenAmount := len(col[0][i])
		lencount[lenAmount]++
	}
	fmt.Println(lencount)
}
