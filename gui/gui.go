// Package gui poskytuje jednoduché konzolové uživatelské rozhraní.
package gui

import (
	"bufio"
	"fmt"
	"os"
)

// AskInput vypíše výzvu a načte uživatelský vstup.
func AskInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1] // odstraníme \n
}
