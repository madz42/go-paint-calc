package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(ask string) string {
	fmt.Print(ask, ":")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func GetDataInt(ask string) (int64, error) {
	return strconv.ParseInt(getInput(ask), 10, 64) 
}

func GetDataFloat(ask string) (float64, error) {
	return strconv.ParseFloat(getInput(ask), 64)
}

func GetDataString(ask string) (string, error) {
	val := getInput(ask)
	if len(val) == 0 {
		return "", errors.New("empty input")
	}
	return strings.ToLower(val), nil
}