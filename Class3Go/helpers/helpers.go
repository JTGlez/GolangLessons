package helpers

import (
	"bufio"
	"fmt"
	"os"
)

// ReadArgsFromStdin reads arguments from standard input until an empty line is entered
func ReadArgsFromStdin() []string {
	var result []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter arguments (press enter to stop):")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}
		result = append(result, input)
	}

	return result
}
