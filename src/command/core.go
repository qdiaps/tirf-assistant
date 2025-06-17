package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func HandleUserInput() {
	fmt.Println("Hello! Enter your command.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input := readLine(reader)

		value, exists := Commands[input]
		if exists == false {
			fmt.Println("Command not found! Enter 'help' to show all commands.")
		} else {
			value.Func()
		}
	}
}

func readLine(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return strings.TrimSpace(input)
}
