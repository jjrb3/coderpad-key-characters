// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Tic-Tac-Toe.
	var ttt [3][3]string

	// Print a draw table.
	fmt.Printf("Welcome to tic tac toe!\n\n")
	drawTable(ttt)

	for {
		fmt.Printf("Write your next move and press enter: ")

		// Reader.
		r := bufio.NewReader(os.Stdin)
		// Enter.
		e, _ := r.ReadString('\n')
		// Remove \n and obtain the selection.
		s := strings.TrimRight(e, "\r\n")
		fmt.Printf("\n\n")

		// Validate input.
		n1, n2, err := validateInput(s, ttt)

		// If exist error again try to write.
		if err {
			fmt.Printf("\n\n")
			continue
		}

		// Set user position and draw table with message [Your move]
		ttt[n1-1][n2-1] = "X"
		drawTable(ttt)
		fmt.Printf("[Your move]")
		fmt.Printf("\n\n")

		// We validate if with it move win.
		if sg := validateWin(ttt); sg {
			break
		}

		// Write the computer move.
		writeComputer(&ttt)
		drawTable(ttt)
		fmt.Printf("[computer move]")

		// We validate if with it move win.
		if sg := validateWin(ttt); sg {
			break
		}

		fmt.Printf("\n\n")
	}
}

func validateWin(ttt [3][3]string) bool {
	// Validate horizontal or vertical
	for i := 0; i < 3; i++ {
		yv := ""
		xv := ""
		for j := 0; j < 3; j++ {
			// Save Y moved
			yv += ttt[i][j]
			// Save X moved
			xv += ttt[j][i]
		}

		// If computer win by X or Y
		if yv == "OOO" || xv == "OOO" {
			fmt.Printf("The computer win!")
			return true
		}

		// If user win by X or Y
		if yv == "XXX" || xv == "XXX" {
			fmt.Printf("You win!")
			return true
		}
	}

	// Diagonal validation
	ld := ttt[0][0] + ttt[1][1] + ttt[2][2]
	rd := ttt[2][0] + ttt[1][1] + ttt[0][2]

	// If computer win by left diagonal or right.
	if ld == "OOO" || rd == "OOO" {
		fmt.Printf("The computer win!")
		return true
	}

	// If user win by left diagonal or right.
	if ld == "XXX" || rd == "XXX" {
		fmt.Printf("You win!")
		return true
	}

	return false
}

func writeComputer(ttt *[3][3]string) {
	for {
		// Random move to computer between 1-3.
		rx := rand.Intn(3)
		ry := rand.Intn(3)

		// If the position is not marked then the computer marke this.
		if ttt[ry][rx] != "X" && ttt[ry][rx] != "O" {
			ttt[ry][rx] = "O"
			return
		}
	}
}

func validateInput(s string, ttt [3][3]string) (int, int, bool) {
	n := strings.Split(s, ",")
	n1, err1 := strconv.Atoi(n[0])
	n2, err2 := strconv.Atoi(n[1])

	// Validate if the param are number
	if err1 != nil || err2 != nil {
		fmt.Print("Ops!, the value registered is incorrect. Try again!")
		return 0, 0, true
	}

	// Validate range of number
	if (n1 < 1 || n1 > 3) ||
		(n2 < 1 || n2 > 3) {
		fmt.Print("Ops!, that cell doesn't exists. Try again!")
		return 0, 0, true
	}

	// Validate marked positions.
	if ttt[n1-1][n2-1] == "O" {
		fmt.Print("Ops!, the position is marked. Try again!")
		return 0, 0, true
	}

	return n1, n2, false
}

func drawTable(ttt [3][3]string) {
	fmt.Print("+---+---+---+\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Position.
			if ttt[i][j] != "" {
				fmt.Printf("| %s ", ttt[i][j])
			} else {
				fmt.Print("|   ")
			}
		}
		fmt.Print("|\n+---+---+---+\n")
	}
}
