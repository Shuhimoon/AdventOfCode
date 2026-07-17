package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"fmt"
)

type Config struct {
	FilePath string `yaml:"file_path"`
	data	 string
	paths	 []string
}

var currentConfig Config

func init() {
	loadConfig()
}

func loadConfig() {
	paths := []string{"config.yaml", "config/config.yaml", "../config.yaml"}
	loaded := false

	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err == nil {
			if err := yaml.Unmarshal(data, &currentConfig); err == nil {
				fmt.Println("成功載入設定檔案: ", p)
				loaded = true
				return
			}
		}

	}

	if loaded == false {
		panic("嚴重錯誤：找不到 config.yaml 或 YAML 格式錯誤！請確認檔案存在且格式正確。")
	}
}


func GetConfig() Config {
	return currentConfig
}
