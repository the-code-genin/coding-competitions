package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TestSetHeader struct {
	m, n, p int
}

var (
	outputStarted bool
	scanner       *bufio.Scanner
)

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

func readNextTestSetHeader(scanner *bufio.Scanner) (*TestSetHeader, error) {
	rawData, err := readNextLine(scanner)
	if err != nil {
		return nil, err
	}

	data := strings.Split(rawData, " ")
	if len(data) != 2 {
		return nil, fmt.Errorf("invalid test set 1")
	}

	m, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 2")
	}

	n, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 3")
	}

	p, err := strconv.Atoi(data[2])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 3")
	}

	return &TestSetHeader{m, n, p}, nil
}

func readParticipantResult(scanner *bufio.Scanner) ([]int, error) {
	rawData, err := readNextLine(scanner)
	if err != nil {
		return nil, err
	}

	data := strings.Split(rawData, " ")
	if len(data) != 2 {
		return nil, fmt.Errorf("invalid test set 1")
	}

	buffer := make([]int, 0)
	for i := 0; i < len(data); i++ {
		res, err := strconv.Atoi(data[i])
		if err != nil {
			return nil, fmt.Errorf("invalid test set 2")
		}

		buffer = append(buffer, res)
	}

	return buffer, nil
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
	scanner = bufio.NewScanner(os.Stdin)
}

func main() {
	// Get number of test sets
	tsLen, err := readLenOfTestSet(scanner)
	if err != nil {
		panic(err)
	}

	// Parse the test sets
	for i := 0; i < tsLen; i++ {
		header, err := readNextTestSetHeader(scanner)
		if err != nil {
			panic(err)
		}

		allRes := make([][]int, 0)
		for j := 0; j < header.m; j++ {
			participantRes, err := readParticipantResult(scanner)
			if err != nil {
				panic(err)
			}

			if len(participantRes) != header.n {
				panic(fmt.Errorf("error in parsing participant rest for tc %v at m %v", i+1, j+1))
			}

			allRes = append(allRes, participantRes)
		}

		reqSteps := 0
		for j := 0; j < header.n; j++ {
			focusPoint := allRes[header.p-1][j]

			maxPoint := focusPoint
			for k := 0; k < header.m; k++ {
				compPoint := allRes[k][j]
				if compPoint > maxPoint {
					maxPoint = compPoint
				}
			}

			if maxPoint > focusPoint {
				reqSteps += maxPoint
			}
		}

		// Opening lines
		printLine(fmt.Sprintf("Case #%d: %d", i+1, reqSteps))
	}
}
