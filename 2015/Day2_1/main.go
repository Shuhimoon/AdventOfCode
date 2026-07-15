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
	TotalArea int
	SquareArea int
	PlusArea int
	Filepaths string
	Data string
	Listdata []string
}

func main() {
	res := Result{}
	res.Filepaths = config.GetConfig().FilePath
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
		res.PlusArea += l * w
		res.SquareArea += 2*l*w + 2*w*h + 2*h*l

	}

	res.TotalArea = res.PlusArea + res.SquareArea
	fmt.Println("Total area: ", res.TotalArea)
}
