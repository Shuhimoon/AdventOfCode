package main

import (
	"fmt"
	"errors"
	"os"
	"strings"
)

type FileInfo struct {
	FilePath string
	LeftCount int
	RightCount int
	Answer int
}


func checkFile(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error:", err)
		}
		return false
	}
	return true
}

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(data)
}
// 讀取.env
func loadEnvConfig() struct{ FilePath string } {
	var cfg struct{ FilePath string }

	checkFile(".env")
	data := ReadFile(".env")

	lines := strings.Split(data, "\n")
	for _,line := range lines{
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line , "=" ,2)
		if len(parts)==2{
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "FILE_PATH" {
				cfg.FilePath = value
				break
			}
		}
	}

	return cfg
}

func main() {
	config := loadEnvConfig()
	checkFile(config.FilePath)
	data := ReadFile(config.FilePath)
	left, right, ans := 0, 0, 0
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		left += strings.Count(line, "(")
		right += strings.Count(line, ")")
	}
	ans = left - right

	result := FileInfo{
		FilePath: config.FilePath,
		LeftCount: left,
		RightCount: right,
		Answer: ans,
	}

	fmt.Printf("檔案: %s\n", result.FilePath)
	fmt.Printf("左括號: %d\n", result.LeftCount)
	fmt.Printf("右括號: %d\n", result.RightCount)
	fmt.Printf("Answer: %d\n", result.Answer)
}
