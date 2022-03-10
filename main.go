package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mockSolvableResult int

type Case struct {
	Pairs    []Pair
	Solution string
}

type Pair struct {
	A, B string
}

func NewCase() *Case {
	return &Case{}
}

func NewPair(a, b string) *Pair {
	return &Pair{a, b}
}

func (c *Case) AddPair(a, b string) {
	c.Pairs = append(c.Pairs, *NewPair(a, b))
}

func (c *Case) ToSolutionString(caseCount int) string {
	return fmt.Sprintf("Case %d: %s", caseCount, c.Solution)
}

func (c *Case) IsSolvable() bool {
	mockSolvableResult++
	return mockSolvableResult%2 == 0
}

func (c *Case) Solve() {
	if c.IsSolvable() {
		var solution string
		// Solving case here
		solution = "mocksolution"

		c.Solution = solution
	} else {
		c.Solution = "IMPOSSIBLE"
	}
}

/*
## Emil's Puzzle
- Read data from file;
- Split data into groups of cases:
-- A case contains an init (int), pairs (two space-separated lowercase alphabetic strings) and three sequences (a
   concatenated string from A' pairs, a concatenated string from B' pairs and a decoded sequence);
- For each case:
-- Verify if case is solvable:
--- If no, add into case.solution a string indicating that this case is impossible to solve (Ex: "Case n: IMPOSSIBLE");
--- If yes, decode a sequence and add result into case.solution (Ex: "Case n: abcdefgh");
-- Print solution;
*/

func main() {
	executeFuncWithTimeTrack("Emil puzzle", problemB)
}

func readDataFromFile(filename string) (result []string) {
	inputFile := getFile(filename)
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)

	for {
		line, _, err := reader.ReadLine()
		if line == nil {
			break
		}
		if err != nil {
			panic(err)
		}
		result = append(result, string(line))
	}

	return
}

func getFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func problemB() {
	inputData := readDataFromFile("sample.in")

	emilPuzzle(inputData)
}

func emilPuzzle(inputData []string) {
	var (
		cases []Case
	)

	for _, data := range inputData {
		_, err := strconv.ParseInt(data, 10, 64)
		if err == nil {
			cases = append(cases, *NewCase())
			continue
		}
		var splitData []string
		splitData = strings.Split(data, " ")
		cases[len(cases)-1].AddPair(splitData[0], splitData[1])
	}

	for i, c := range cases {
		c.Solve()
		fmt.Println(c.ToSolutionString(i + 1))
	}
}
