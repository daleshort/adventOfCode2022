
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	filePath := "day8.txt"
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("cannot open file")
	}

	dataRow := strings.Split(string(data), "\n")
	dataMatrix := make([][]int, len(dataRow))

	for index, row := range dataRow {
		dataMatrix[index] = castStringSliceToIntSlice(strings.Split(row, ""))
	}

	locations := make(map[string]struct{})

	//go through rows
	for row, _ := range dataMatrix {
		cols := getVisibleLocations(getRow(dataMatrix, row))
		for _, col := range cols {
			addToLocationsIfUnique(locations, row, col)
		}
	}

	// go through cols
	for col, _ := range dataMatrix[0] {
		rows := getVisibleLocations(getCol(dataMatrix, col))
		for _, row := range rows {
			addToLocationsIfUnique(locations, row, col)
		}
	}

	fmt.Println("number trees", len(locations))

	//part 2
	maxScore := 0
	for row, _ := range dataMatrix {
		for col, _ := range dataMatrix[0] {
			go processPoint(row, col, dataMatrix, &maxScore)
		}
	}

	fmt.Println("max score", maxScore)

	elapsed := time.Since(start)
	log.Printf("execution took %s", elapsed)
}

func processPoint(row int, col int, dataMatrix [][]int, maxScore *int) {
	curRow := getRow(dataMatrix, row)
	curCol := getCol(dataMatrix, col)
	score := getViewScore(curRow, col) * getViewScore(curCol, row)
	*maxScore = max(*maxScore, score)
}

func getViewScore(list []int, start int) int {

	if start == len(list)-1 || start == 0 {
		return 0
	}
	startHeight := list[start]
	score := 1

	//starting at start and going right
	for i := start + 1; i < len(list); i++ {
		currentHeight := list[i]
		if currentHeight >= startHeight || i == len(list)-1 {
			score *= i - start

			break
		}
	}

	//start at start and go left
	for i := start - 1; i >= 0; i-- {
		currentHeight := list[i]
		if currentHeight >= startHeight || i == 0 {
			score *= start - i
			break
		}
	}

	return score

}

func addToLocationsIfUnique(locations map[string]struct{}, row int, col int) {
	stringKey := makeStringKey(row, col)
	if _, ok := locations["foo"]; !ok {
		locations[stringKey] = struct{}{}
	}
}

func makeStringKey(row int, col int) string {
	return fmt.Sprintf("%v-%v", row, col)
}

func castStringSliceToIntSlice(data []string) []int {

	returnSlice := make([]int, len(data))
	for index, val := range data {
		var err error
		returnSlice[index], err = strconv.Atoi(val)
		if err != nil {
			log.Fatal("error parsing int")
		}
	}
	return returnSlice
}

func getRow(dataMatrix [][]int, row int) []int {

	return dataMatrix[row]
}

func getCol(dataMatrix [][]int, col int) []int {
	var returnSlice []int

	for _, row := range dataMatrix {
		returnSlice = append(returnSlice, row[col])
	}
	return returnSlice
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}

func getVisibleLocations(line []int) []int {
	//from left
	var locations []int
	lastHeight := 0

	for i := 0; i < len(line); i++ {
		if i == 0 {
			locations = append(locations, i)
			lastHeight = max(line[i], lastHeight)
		} else if line[i] > lastHeight {
			locations = append(locations, i)
			lastHeight = max(line[i], lastHeight)
		}
	}

	lastHeight = 0

	for i := len(line) - 1; i >= 0; i-- {
		if i == len(line)-1 {
			locations = append(locations, i)
			lastHeight = max(line[i], lastHeight)
		} else if line[i] > lastHeight {
			locations = append(locations, i)
			lastHeight = max(line[i], lastHeight)
		}
	}
	return unique(locations)

}
