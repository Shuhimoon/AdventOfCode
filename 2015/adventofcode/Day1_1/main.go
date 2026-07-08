package main

import (
	"fmt"
	"errors"
	"os"
	"strings"
	// 從config讀取檔案路徑
	"Day1_1/config"
)

type FileInfo struct {
	FilePath string
	LeftCount int
	RightCount int
	Answer int
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

	left := strings.Count(data, "(")
	right := strings.Count(data, ")")
	ans := left - right

	result := FileInfo{
		FilePath: filePath,
		LeftCount: left,
		RightCount: right,
		Answer: ans,
	}

	fmt.Printf("檔案: %s\n", result.FilePath)
	fmt.Printf("左括號: %d\n", result.LeftCount)
	fmt.Printf("右括號: %d\n", result.RightCount)
	fmt.Printf("Answer: %d\n", result.Answer)
}
