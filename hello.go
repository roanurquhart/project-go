package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Start a TicTacToe Game?")
	fmt.Println("Press Y to start or N to quit:")
	var response string
	fmt.Scanln(&response)
	var gameBoard = board{"   ", "   ", "   "}
	printTicTacToe(gameBoard)
	for response == "Y" {
		response = playerOneTurn(gameBoard)
	}
	fmt.Println("Thanks for playing!")
}

func printTicTacToe(tictactoe board) {
	fmt.Println("|", string(tictactoe.first[0]), "|", string(tictactoe.first[1]), "|", string(tictactoe.first[2]), "|")
	fmt.Println("|", string(tictactoe.second[0]), "|", string(tictactoe.second[1]), "|", string(tictactoe.second[2]), "|")
	fmt.Println("|", string(tictactoe.third[0]), "|", string(tictactoe.third[1]), "|", string(tictactoe.third[2]), "|")
}

func playerOneTurn(gameBoard board) string {
	fmt.Println("Player One:\n Enter your move with row number followed by column number.\n Example: 1,2")
	var response string
	fmt.Scanln(&response)
	row, column, isQuitting := parseInput(response)

	if isQuitting {
		return "N"
	}

	for row != "1" && row != "2" && row != "3" && column < 1 && column > 3 {
		fmt.Println("Invalid input. Please enter again")
		fmt.Scanln(&response)
		row, column, isQuitting = parseInput(response)

		if isQuitting {
			return "N"
		}
	}
	fmt.Println()

	switch row {
	case "1":
		gameBoard.first = replaceAtIndex(gameBoard.first, 'X', column-1)
	case "2":
		gameBoard.second = replaceAtIndex(gameBoard.second, 'X', column-1)
	case "3":
		gameBoard.third = replaceAtIndex(gameBoard.third, 'X', column-1)
	}
	printTicTacToe(gameBoard)
	return "Y"
}

func parseInput(input string) (row string, column int, quit bool) {
	quit = false
	splitString := strings.Split(input, ",")
	row = splitString[0]
	if row == "q" || row == "Q" {
		fmt.Println("Quitting...")
		column = 0
		quit = true
		return
	}
	if len(splitString) < 2 {
		fmt.Print("Should get invalid message")
		column = 0
		return
	}
	column, _ = strconv.Atoi(splitString[1])
	return
}

// Thanks to OneOfOne from this StackOverflow post:
// https://stackoverflow.com/questions/24893624/how-to-replace-a-letter-at-a-specific-index-in-a-string-in-go
func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

type board struct {
	first  string
	second string
	third  string
}
