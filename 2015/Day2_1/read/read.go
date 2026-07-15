package read
import (
	"fmt"
	"os"
	"errors"
)

type Input struct {
	Data string
	Err  error
}

func checkFile(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err,os.ErrNotExist){
			panic("file not found")
		}else{
			panic(err)
		}
	}
	return nil
}

func ReadInput(filePath string) (string, error) {
	checkFile(filePath)
	data,err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error Read File ：",err)
		return "", err
	}
	return string(data), nil
}
