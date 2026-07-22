package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"Day2_1/config"
	"Day2_1/read"
)

type Result struct {
	TotalLength int
	RibbonLength int
	PlusLength int
	Filepaths string
	Data string
	Listdata []string
}


var res Result = Result{
	Filepaths: config.GetConfig().FilePath,
}


func main() {
	res.Data, _ = read.ReadInput(res.Filepaths)
	// Split the data into lines
	res.Listdata = strings.Split(res.Data, "\n")
	for _, line := range res.Listdata {
		square := strings.Split(line, "x")
		// Check if the line is a valid square
		if len(square) != 3 {
			continue
		}
		// Sort the lengths
		sortlen := make([]int, 3)
		for i, s := range square {
			sortlen[i], _ = strconv.Atoi(s)
		}
		slices.Sort(sortlen)
		// Convert the lengths to integers
		l, w, h := sortlen[0], sortlen[1], sortlen[2]
		res.PlusLength += l*2 + w*2
		res.RibbonLength += w*h*l

	}

	res.TotalLength = res.PlusLength + res.RibbonLength
	fmt.Println("Total length: ", res.TotalLength)
}
