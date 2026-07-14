package main

import (
	"fmt"
	"errors"
	"os"
	// 從config讀取檔案路徑
	"Day1_1/config"
)

type FileInfo struct {
	FilePath string
	LeftCount int
	RightCount int
	Answer int
	CountTime int
}


func checkFile(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			panic("File does not exist")
		} else {
			panic("Error reading file")
		}
	}
}

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(data)
}

func main() {
	// 讀取config檔案路徑
	filePath := config.GetConfig().FilePath
	// 如果輸入指令時 有帶檔案路徑
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}
	// 檢查檔案是否存在
	checkFile(filePath)
	data := ReadFile(filePath)

	result := FileInfo{
		FilePath: filePath,
		LeftCount: 0,
		RightCount: 0,
		Answer: 0,
		CountTime: 0,
	}

	for i,char := range data {
		if char == '(' {
			result.Answer++
			result.LeftCount++
			// fmt.Printf("位置 %d: '%c' → ans = %d (左括號)\n", i, char, result.Answer)
		} else if char == ')' {
			result.Answer--
			result.RightCount++
			// fmt.Printf("位置 %d: '%c' → ans = %d (右括號)\n", i, char, result.Answer)
		}


		if result.Answer == -1 {
			fmt.Printf("檔案: %s\n", result.FilePath)
			fmt.Printf("左括號: %d\n", result.LeftCount)
			fmt.Printf("右括號: %d\n", result.RightCount)
			fmt.Printf("Answer: %d\n", result.Answer)
			fmt.Printf("CountTime: %d\n", i+1)
			result.CountTime = i + 1
			break
		}
	}
}
