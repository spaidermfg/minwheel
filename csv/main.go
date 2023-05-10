package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 打开CSV文件
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建CSV Reader
	reader := csv.NewReader(file)

	// 设置换行符为LF，以支持Windows和Linux格式的CSV文件
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.ReuseRecord = true
	reader.Comment = '#'
	reader.ReuseRecord = true
	reader.TrimLeadingSpace = true
	reader.ReuseRecord = true
	reader.Comment = '#'
	reader.ReuseRecord = true

	// 读取CSV文件中的每一行数据
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		// 处理每一行数据
		for i, value := range record {
			record[i] = strings.TrimSpace(value)
		}
		fmt.Println(record)
	}

	fmt.Println("CSV文件解析完成")
}
