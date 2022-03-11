package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	var (
		isSolvable bool
		matchCount int
	)

	for _, pair := range c.Pairs {
		if matchCount > 1 {
			return false
		}
		if strings.HasPrefix(pair.A, pair.B) || strings.HasPrefix(pair.B, pair.A) {
			isSolvable = true
			matchCount++
		}
	}

	return isSolvable
}

func (c *Case) Solve() {
	if c.IsSolvable() {
		c.Decode("", "", c.Pairs)
	} else {
		c.Solution = "IMPOSSIBLE"
	}
}

func (c *Case) Decode(sequenceA, sequenceB string, pairs []Pair) (string, string, []Pair) {
	var partialA, partialB string
	for index, pair := range pairs {
		if strings.EqualFold(pair.A, pair.B) {
			remainingPairs := removeCurrentIndex(index, pairs)
			return c.Decode(sequenceA, sequenceB, remainingPairs)
		}
		partialA = sequenceA + pair.A
		partialB = sequenceB + pair.B
		if strings.HasPrefix(partialA, partialB) || strings.HasPrefix(partialB, partialA) {
			if len(partialA) > len(partialB) {
				sequenceA = partialA
				sequenceB = partialB
				remainingPairs := removeCurrentIndex(index, pairs)
				return c.Decode(sequenceA, sequenceB, remainingPairs)
			} else {
				sequenceA = partialA
				sequenceB = partialB
				remainingPairs := removeCurrentIndex(index, pairs)
				return c.Decode(sequenceA, sequenceB, remainingPairs)
			}
		}
	}

	if sequenceA == sequenceB {
		c.Solution = sequenceA
	} else {
		c.Solution = "IMPOSSIBLE"
	}
	return "", "", nil
}

func removeCurrentIndex(index int, pairs []Pair) (result []Pair) {
	for i, pair := range pairs {
		if i == index {
			continue
		}
		result = append(result, pair)
	}
	return
}

func main() {
	problemB()
}

func readData() (inputData []string) {
	var count int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		inputData = append(inputData, scanner.Text())
		if len(inputData) > 0 {
			log.Println(inputData[count])
			count++
		}
	}

	return
}

func readDataFromFile() (inputData []string) {
	inputFile := getFile("sample-01.in")
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
		inputData = append(inputData, string(line))
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
	//inputData := readDataFromFile()
	inputData := readData()

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
