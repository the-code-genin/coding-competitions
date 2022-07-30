package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TestSet struct {
	rows, cols int
}

var outputStarted bool

func readNextLine(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func readLenOfTestSet(scanner *bufio.Scanner) (int, error) {
	data, err := readNextLine(scanner)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(data)
}

func readNextTestSet(scanner *bufio.Scanner) (*TestSet, error) {
	rawData, err := readNextLine(scanner)
	if err != nil {
		return nil, err
	}

	data := strings.Split(rawData, " ")
	if len(data) != 2 {
		return nil, fmt.Errorf("invalid test set 1")
	}

	row, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 2")
	}

	col, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 3")
	}

	return &TestSet{row, col}, nil
}

func genFrame(set *TestSet, t string) string {
	data := make([]string, 0)
	if t == "start" {
		data = append(data, ".")
	} else {
		data = append(data, "+")
	}

	for i := 0; i < set.cols; i++ {
		if i != 0 || t != "start" {
			data = append(data, "-+")
		} else {
			data = append(data, ".+")
		}
	}

	return strings.Join(data, "")
}

func genRow(row, cols int) string {
	data := make([]string, 0)
	if row == 0 {
		data = append(data, ".")
	} else {
		data = append(data, "|")
	}

	for i := 0; i < cols; i++ {
		data = append(data, ".|")
	}

	return strings.Join(data, "")
}

func printLine(data string) {
	// Should the line break be printed
	if outputStarted {
		fmt.Print("\n")
	} else {
		outputStarted = true
	}

	fmt.Print(data)
}

func init() {
	outputStarted = false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Get number of test sets
	tsLen, err := readLenOfTestSet(scanner)
	if err != nil {
		panic(err)
	}

	// Parse the test sets
	for i := 0; i < tsLen; i++ {
		set, err := readNextTestSet(scanner)
		if err != nil {
			panic(err)
		}

		// Opening lines
		printLine(fmt.Sprintf("Case #%d:", i + 1))
		printLine(genFrame(set, "start"))

		for j := 0; j < set.rows; j++ {
			printLine(genRow(j, set.cols))
			printLine(genFrame(set, "end"))
		}
	}
}
