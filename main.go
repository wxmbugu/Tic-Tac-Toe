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
	fmt.Println("| " + board[0] + "| " + board[1] + "| " + board[2] + "|")
	fmt.Println("| " + board[3] + "| " + board[4] + "| " + board[5] + "|")
	fmt.Println("| " + board[6] + "| " + board[7] + "| " + board[8] + "|")
	fmt.Println("-------")
	return board
}
func boardprinter() {
	fmt.Println("Here is a board for you to make a mark using the specified letters:")
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
func placepiece(board *[9]string, position string, player string) {
	var symbol string
	winner, ok := win(board, player)
	if ok == true {
		fmt.Println("Winner is: ", winner)
		os.Exit(1)
	}
	if player == "cpu" {
		symbol = "X"
		fmt.Println("Player X(comp) is playing... ")
		cpuPos := randomStr(1)
		position = cpuPos
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
	if board[0] != "" && board[1] != "" && board[2] != "" && board[3] != "" && board[4] != "" && board[5] != "" && board[6] != "" && board[7] != "" && board[8] != "" {
		winner, ok := win(board, player)
		if ok == true {
			fmt.Println("Winner is: ", winner)
			os.Exit(1)
		} else {
			fmt.Println("It's a draw!")
			bmarker(*board)
			os.Exit(1)
		}
	}
	switch position {
	case "a":
		if board[0] == "" {
			board[0] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "b":
		if board[1] == "" {
			board[1] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "c":
		if board[2] == "" {
			board[2] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "d":
		if board[3] == "" {
			board[3] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "e":
		if board[4] == "" {
			board[4] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "f":
		if board[5] == "" {
			board[5] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "g":
		if board[6] == "" {
			board[6] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "h":
		if board[7] == "" {
			board[7] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	case "i":
		if board[8] == "" {
			board[8] = symbol
		} else {
			fmt.Println("Position is filled , enter another position")
			placepiece(board, position, player)
		}
	default:
		fmt.Println("***WARNING***")
		fmt.Println("That was an invalid entry. Please try again player:", symbol)
		placepiece(board, position, player)
		break
	}
}
func main() {
	var ttt [9]string
	var position string
	var winner string
	for {
		boardprinter()
		cpuPos := randomStr(1)
		placepiece(&ttt, position, "player")
		placepiece(&ttt, cpuPos, "cpu")
		bmarker(ttt)
		winner, ok := win(&ttt, winner)
		if ok == true {
			fmt.Println("Winner is: ", winner)
			break
		}
	}
}

//function to get random letters from (a-i)
//so as to randomise game play with computer instead of using an ai
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

func win(ttt *[9]string, winner string) (string, bool) {
	if ttt[0] == ttt[1] {
		if ttt[1] == ttt[2] {
			if ttt[2] != "" {
				winner = ttt[2]
				return winner, true
			}
		}
	}
	//Wins by matching the middle  row
	if ttt[3] == ttt[4] {
		if ttt[4] == ttt[5] {
			if ttt[3] != "" {
				return winner, true
			}

		}
	}
	//Wins by matching the last row
	if ttt[6] == ttt[7] {
		if ttt[7] == ttt[8] {
			if ttt[6] != "" {
				winner = ttt[6]
				return winner, true
			}
		}
	}
	//Wins by matching the left diagnol row
	if ttt[0] == ttt[4] {
		if ttt[4] == ttt[8] {
			if ttt[0] != "" {
				winner = ttt[0]
				return winner, true
			}
		}
	}
	//Wins by matching the right diagnol row
	if ttt[2] == ttt[4] {
		if ttt[4] == ttt[6] {
			if ttt[4] != "" {
				winner = ttt[4]
				return winner, true
			}
		}
	}
	//wins by matching the first column
	if ttt[0] == ttt[3] {
		if ttt[3] == ttt[6] {
			if ttt[3] != "" {
				winner = ttt[3]
				return winner, true
			}
		}
	} //wins by matching the middle column
	if ttt[1] == ttt[4] {
		if ttt[4] == ttt[7] {
			if ttt[7] != "" {
				winner = ttt[7]
				return winner, true
			}
		}
	}
	//wins by matching the last column
	if ttt[2] == ttt[5] {
		if ttt[5] == ttt[8] {
			if ttt[5] != "" {
				winner = ttt[5]
				return winner, true
			}
		}
	}
	return "", false
}
