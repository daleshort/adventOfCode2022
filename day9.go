package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type location struct {
	x int
	y int
}

type ropeEnd struct {
	loc location
}

type gameBoard struct {
	head          ropeEnd
	tail          ropeEnd
	tailLocations map[string]struct{}
}

func (end *ropeEnd) step(dir string) {

	switch dir {
	case "R":
		end.loc.x += 1
	case "L":
		end.loc.x -= 1
	case "U":
		end.loc.y += 1
	case "D":
		end.loc.y -= 1
	}
}

func (gameBoard *gameBoard) processStep(dir string) {
	fmt.Println("------")
	fmt.Println("new step of dir", dir)

	gameBoard.head.step(dir)

	if gameBoard.isTailMoveNeeded() {
		fmt.Println("\t\t00000000000")
		//gameBoard.printBoard(true)
		gameBoard.moveTail()
		fmt.Println("\t\t00000000000")
	}
	//gameBoard.printBoard(false)
	gameBoard.registerTailLocation()
}

func (g *gameBoard) registerTailLocation() {
	key := makeStringKey(g.tail.loc.x, g.tail.loc.y)

	_, ok := g.tailLocations[key]
	if !ok {
		g.tailLocations[key] = struct{}{}
	}

}

func makeStringKey(row int, col int) string {
	return fmt.Sprintf("%v-%v", row, col)
}

func (g *gameBoard) isTailMoveNeeded() bool {
	distX := Abs(g.calculateHeadToTail("x"))
	distY := Abs(g.calculateHeadToTail("y"))

	if distX > 1 || distY > 1 {
		fmt.Println("move needed")
		return true
	}
	fmt.Println("tail good")
	return false
}

func (g *gameBoard) calculateHeadToTail(dir string) int {
	if dir == "x" {
		return g.head.loc.x - g.tail.loc.x
	} else if dir == "y" {
		return g.head.loc.y - g.tail.loc.y
	}
	return 0
}

func (g *gameBoard) moveTail() {
	deltaX, deltaY := g.calculateHeadToTail("x"), g.calculateHeadToTail("y")
	fmt.Println("deltaX:", deltaX, "deltaY", deltaY)
	if deltaY == 0 {
		if deltaX < 0 {
			g.tail.step("L")
		} else {
			g.tail.step("R")
		}
	} else if deltaX == 0 {
		if deltaY < 0 {
			g.tail.step("D")
		} else {
			g.tail.step("U")
		}
	} else if deltaY == 2{
		if deltaX < 0 {
			g.tail.step("L")
			g.tail.step("U")
		} else {
			g.tail.step("R")
			g.tail.step("U")
		}
	} else if deltaY == -2 {
		if deltaX < 0 {
			g.tail.step("L")
			g.tail.step("D")
		} else {
			g.tail.step("R")
			g.tail.step("D")
		}
	} else if deltaX == -2 {
		if deltaY < 0 {
			g.tail.step("L")
			g.tail.step("D")
		} else {
			g.tail.step("U")
			g.tail.step("L")
		}
	} else if deltaX == 2 {
		if deltaY < 0 {
			g.tail.step("R")
			g.tail.step("D")
		} else {
			g.tail.step("U")
			g.tail.step("R")
		}
	}

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (gameBoard *gameBoard) printBoard(useTab bool) {
	if useTab {
		fmt.Printf("\t\t")
	}
	fmt.Printf("head| %v , %v \n", gameBoard.head.loc.x, gameBoard.head.loc.y)
	if useTab {
		fmt.Printf("\t\t")
	}
	fmt.Printf("tail| %v , %v \n", gameBoard.tail.loc.x, gameBoard.tail.loc.y)

	maxX := int(math.Max(float64(gameBoard.head.loc.x), float64(gameBoard.tail.loc.x)))
	maxY := int(math.Max(float64(gameBoard.head.loc.y), float64(gameBoard.tail.loc.y)))
	if useTab {
		fmt.Printf("\t\t")
	}
	fmt.Printf("Max X %v Max Y %v \n", maxX, maxY)
	boardArt := make([]string, 0)
	var rowString string

	for i := 0; i <= maxX; i++ {
		rowString += "."
	}
	for i := 0; i <= maxY; i++ {
		boardArt = append(boardArt, rowString)
	}
	boardArt[gameBoard.tail.loc.y] = replaceCharacterAtIndex(boardArt[gameBoard.tail.loc.y], rune('T'), gameBoard.tail.loc.x)
	boardArt[gameBoard.head.loc.y] = replaceCharacterAtIndex(boardArt[gameBoard.head.loc.y], rune('H'), gameBoard.head.loc.x)

	for i := len(boardArt) - 1; i >= 0; i-- {
		if useTab {
			fmt.Printf("\t\t")
		}
		fmt.Printf("%v \n", boardArt[i])

	}

}

func replaceCharacterAtIndex(s string, replacement rune, index int) string {
	out := []rune(s)
	out[index] = replacement
	return string(out)
}

func (gameBoard *gameBoard) initBoard() {
	tInitial := location{0, 0}
	hInitial := location{0, 0}

	tail := ropeEnd{tInitial}
	head := ropeEnd{hInitial}

	gameBoard.head = head
	gameBoard.tail = tail

	gameBoard.tailLocations = make(map[string]struct{})
}

func main() {
	filePath := "day9.txt"
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("cannot open file")
	}
	var board gameBoard
	board.initBoard()
	dataRow := strings.Split(string(data), "\n")

	for _, row := range dataRow {
		dir, count := parseRow(row)

		for i := 0; i < count; i++ {
			board.processStep(dir)
		}
	}

	fmt.Println("final count of tail locations:", len(board.tailLocations))
}

func parseRow(rowData string) (string, int) {
	splitRow := strings.Split(rowData, " ")
	dir := splitRow[0]
	count, _ := strconv.Atoi(splitRow[1])
	return dir, count
}
