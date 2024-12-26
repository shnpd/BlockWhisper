package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	// 打开 Excel 文件
	filePath := "data/bitcoin_20240911.xlsx" // 替换为你的 Excel 文件路径
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
	// 读取工作表名称
	sheetName := "blockchair_bitcoin_transactions" // 替换为你的工作表名称
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("无法读取工作表: %v", err)
	}

	var amount []string
	for i := 1; i < len(rows); i++ {
		if rows[i][1] == "1" {
			amount = append(amount, rows[i][3])
		}
		if rows[i][2] == "1" {

			amount = append(amount, rows[i][4])
		}
	}

	saveAmount(amount)

}

func saveAmount(data []string) {
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 定义工作表名称
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)

	// 遍历切片，将每个元素写入到 Excel 的一列中
	column := "A" // 要写入的列，例如 A 列
	for i, value := range data {
		cell := fmt.Sprintf("%s%d", column, i+1) // 拼接单元格名称，例如 A1, A2, ...
		if err := f.SetCellValue(sheetName, cell, value); err != nil {
			log.Fatalf("写入单元格 %s 失败: %v", cell, err)
		}
	}
	// 设置工作表为活动状态
	f.SetActiveSheet(index)

	// 保存 Excel 文件
	filePath := "normal_amount.xlsx" // 保存路径
	if err := f.SaveAs(filePath); err != nil {
		log.Fatalf("保存文件失败: %v", err)
	}

	fmt.Printf("数据已成功保存到 %s 的 %s 列\n", filePath, column)

}
