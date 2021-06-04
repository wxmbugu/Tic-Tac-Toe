package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func bmarker(board [9]string) [9]string {
	fmt.Println("The board now looks like this:")
	fmt.Println("-------")
	fmt.Println("|" + board[0] + " | " + board[1] + "| " + board[2] + "|")
	fmt.Println("|" + board[3] + " | " + board[4] + "| " + board[5] + "|")
	fmt.Println("|" + board[6] + " | " + board[7] + "| " + board[8] + "|")
	fmt.Println("-------")
	return board

}
func boardprinter() {
	fmt.Println("Here is a board for you to make a mark:")
	fmt.Println("-------")
	fmt.Println("|a|b|c|")
	fmt.Println("|d|e|f|")
	fmt.Println("|g|h|i|")
	fmt.Println("-------")
}

//Step 1:Start by printing the board
//Step 2:Inputting to board
//step 2.2:Saving into board all entries
//Step 3:Players
//Step 4:Checking rules of the game or winner and draws

func placepiece(board [9]string, position string, player string) {
	var symbol = " "
	if player == "cpu" {
		symbol = "X"
		fmt.Println("Player X(comp) is playing... ")
	}
	if player == "player" {
		symbol = "O"
		fmt.Println("Player O is playing... ")
		fmt.Print("Please enter a letter for the square you want to mark (a through i): \n")
		_, err := fmt.Scanln(&position)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(position)
	}
	switch position {
	case "a":
		board[0] = symbol
		break
	case "b":
		board[1] = symbol
		break
	case "c":
		board[2] = symbol
		break
	case "d":
		board[3] = symbol
		break
	case "e":
		board[4] = symbol
		break
	case "f":
		board[5] = symbol
		break
	case "g":
		board[6] = symbol
		break
	case "h":
		board[7] = symbol
		break
	case "i":
		board[8] = symbol
		break
	default:
		fmt.Println("***WARNING***")
		fmt.Println("That was an invalid entry. Please try again player")
		placepiece(board, position, player)
		break
	}
	bmarker(board)
}

func main() {
	boardprinter()
	var ttt [9]string
	var position string

	//Only lowercase random letters
	for true {
		cpuPos := randomStr(1)
		placepiece(ttt, position, "player")
		placepiece(ttt, cpuPos, "cpu")
	}

}

func randomStr(length int) string {
	rand.Seed(time.Now().Unix())
	charSet := "abcdefghi"
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()

}
