package main

import(
	// "bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"Day5_1/config"
	"Day5_1/read"
)

// Result 結構用來儲存解析後的結果
type Result struct {
	Filepaths string
	Data      string
	List      []string
	Count int
	leftpointer [2]int
	rightpointer [2]int
	area int
}

type DiffMatrix struct {
	diff []uint8
	rows int
	cols int
}


// 設定預設數值
var res Result = Result{
	Filepaths: config.GetConfig().FilePath,
	Data: "",
	Count: 0,
}

// 解析(Sscanf) 拆出座標 ，將左邊座標和右邊座標存入 Result 結構中
func Left_Pointer(line string, action string) {
	x1, y1, x2, y2 := 0, 0, 0, 0
	fmt.Sscanf(line, action+" %d,%d through %d,%d", &x1, &y1, &x2, &y2)

	// 強制保證 left 的座標 ≤ right 的座標
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		res.leftpointer = [2]int{x1, y1}
		res.rightpointer = [2]int{x2, y2}


}

// NewDiffMatrix 用來創建一個新的 DiffMatrix 結構 ， 長寬自動加一 ， 並初始化所有元素為 0
func NewDiffMatrix(rows, cols int) *DiffMatrix{
	return &DiffMatrix{
		diff: make([]uint8, rows*cols),
		rows: rows,
		cols: cols,
	}
}

// Add：翻轉矩形區域（toggle）
func (d *DiffMatrix) Add(x1, y1, x2, y2 int, _ uint8) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			d.diff[y*d.cols+x] ^= 1 // 0↔1
		}
	}
}

// 開燈或關燈 維持
func (d *DiffMatrix) Keep(x1, y1, x2, y2 int, value uint8) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			d.diff[y*d.cols+x] = value
		}
	}
}

func main() {
	res.Data , _ = read.ReadInput(res.Filepaths)
	res.List = strings.Split(res.Data, "\n")
	// 取得區域大小
	if len(os.Args) > 1 {
		res.area, _ = strconv.Atoi(os.Args[1])
	} else {
		res.area = 1000
	}

	// 初始化 DiffMatrix
	dm := NewDiffMatrix(res.area, res.area)
	for _, line := range res.List {
		// 判斷是哪個指令
		switch {
		// HasPrefix 是否有指定開頭
		case strings.HasPrefix(line, "turn off"):
			// 取得指令中的座標
			Left_Pointer(line,"turn off")
			// 將 DiffMatrix 中對應的區域減去 1
			dm.Keep(res.leftpointer[0], res.leftpointer[1], res.rightpointer[0], res.rightpointer[1], 0)

		case strings.HasPrefix(line, "turn on"):
			Left_Pointer(line,"turn on")
			// 將 DiffMatrix 中對應的區域增加 1
			dm.Keep(res.leftpointer[0], res.leftpointer[1], res.rightpointer[0], res.rightpointer[1], 1)

		case strings.HasPrefix(line, "toggle"):
			Left_Pointer(line,"toggle")
			// 將 DiffMatrix 中對應的區域取反
			dm.Add(res.leftpointer[0], res.leftpointer[1], res.rightpointer[0], res.rightpointer[1], 1)

		default:
			continue
		}
	}
	for _, v := range dm.diff {
		if v == 1 {
			res.Count++
		}
	}
	fmt.Println(res.Count)

}
