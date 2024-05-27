package utils

import (
	"bufio"
	"fmt"
	"strings"
)

var BaseURL = "http://localhost:9090"

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
