package validators

import (
	"bufio"
	"fmt"
	"strconv"
)

func IDValidator(val string, reader *bufio.Reader) (int, error) {
	id, err := strconv.Atoi(val)
	if err != nil || id < 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return 0, err
	}

	return id, nil
}

func FundsValidator(val string, reader *bufio.Reader) (float64, error) {
	amount, err := strconv.ParseFloat(val, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return 0, err
	}

	return amount, nil
}

func MoveValidator(val string, reader *bufio.Reader) (int, error) {
	move, err := strconv.Atoi(val)
	if err != nil || move < 0 || move > 2 {
		fmt.Println("Invalid input. Please enter a number between 0 and 2.")
		return 0, err
	}

	return move, nil
}
