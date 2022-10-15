package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TestSetHeader struct {
	rs, rh int
}

type BallPoint struct {
	x, y, t int
}

func (p *BallPoint) Distance() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2))
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

	rs, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 2")
	}

	rh, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, fmt.Errorf("invalid test set 3")
	}

	return &TestSetHeader{rs, rh}, nil
}

func readTeamResult(scanner *bufio.Scanner, team int) ([]*BallPoint, error) {
	rawData, err := readNextLine(scanner)
	if err != nil {
		return nil, err
	}

	n, err := strconv.Atoi(rawData)
	if err != nil {
		return nil, fmt.Errorf("invalid test set 4")
	}

	buffer := make([]*BallPoint, 0)
	for i := 0; i < n; i++ {
		rawData, err = readNextLine(scanner)
		if err != nil {
			return nil, err
		}

		data := strings.Split(rawData, " ")
		if len(data) != 2 {
			return nil, fmt.Errorf("invalid test set 5")
		}

		x, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, fmt.Errorf("invalid test set 6")
		}

		y, err := strconv.Atoi(data[1])
		if err != nil {
			return nil, fmt.Errorf("invalid test set 7")
		}

		buffer = append(buffer, &BallPoint{x, y, team})
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
		tcHeader, err := readNextTestSetHeader(scanner)
		if err != nil {
			panic(err)
		}

		rawTeamResults := make([]*BallPoint, 0)
		for j := 0; j < 2; j++ {
			teamRes, err := readTeamResult(scanner, j)
			if err != nil {
				panic(err)
			}

			rawTeamResults = append(rawTeamResults, teamRes...)
		}

		teamResults := make([]*BallPoint, 0)
		for j := 0; j < len(rawTeamResults); j++ {
			res := rawTeamResults[j]
			if res.Distance() <= float64(tcHeader.rh) + float64(tcHeader.rs) {
				teamResults = append(teamResults, res)
			}
		}

		for j := 0; j < len(teamResults); j++ {
			smallestValue := teamResults[j]
			smallestIndex := j
	
			for k := j; k < len(teamResults); k++ {
				if teamResults[k].Distance() < smallestValue.Distance() {
					smallestValue = teamResults[k]
					smallestIndex = k
				}
			}
	
			if smallestIndex != j {
				oldSmallestValue := teamResults[j]
		
				teamResults[j] = smallestValue
				teamResults[smallestIndex] = oldSmallestValue
			}
		}

		teamPoints := make([]int, 2)
		for j := 0; j < len(teamResults); j++ {
			currentRes := teamResults[j]
			endOfScoring := false

			for k := 0; k < j; k++ {
				prevRes := teamResults[k]
				if prevRes.t != currentRes.t {
					endOfScoring = true
					break
				}
			}

			if !endOfScoring {
				teamPoints[currentRes.t] += 1
			} else {
				break
			}
		}

		// Opening lines
		printLine(fmt.Sprintf("Case #%d: %d %d", i+1, teamPoints[0], teamPoints[1]))
	}
}
