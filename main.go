package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	impossibleSolution = "IMPOSSIBLE"
)

func main() {
	//executeFuncWithTimeTrack("nuno", nuno)
	nuno()
}

func nuno() {
	inputData := getInputData()

	cases := parseInputToCases(inputData)

	for i, currentCase := range cases {
		evaluate(&currentCase)
		fmt.Println(currentCase.ToSolutionString(i + 1))
	}
}

func getInputData() []string {
	scope := os.Getenv("SCOPE")
	if strings.EqualFold(scope, "LOCAL") {
		return readDataFromFile()
	}
	return readDataFromKeyboard()
}

func evaluate(currentCase *Case) {
	if currentCase.IsSolvable() {
		currentCase.Solution = currentCase.Decode()
	} else {
		currentCase.Solution = NewImpossibleSolution()
	}
}

func (c *Case) Decode() string {
	var (
		partialA, partialB string
		matches            []Match
	)

	for index, pair := range c.Pairs {
		if strings.EqualFold(pair.A, pair.B) {
			c.RemovePair(index)
			return c.Decode()
		}

		partialA = strings.ToLower(c.Sequence.A + pair.A)
		partialB = strings.ToLower(c.Sequence.B + pair.B)

		if strings.HasPrefix(partialA, partialB) {
			matches = append(matches, *NewMatch(index, c.Sequence.A+pair.A))
		} else if strings.HasPrefix(partialB, partialA) {
			matches = append(matches, *NewMatch(index, c.Sequence.B+pair.B))
		}
	}

	if matchesShouldBeenSorted(matches) {
		matches = sortMatches(matches)
	}

	if len(matches) > 0 {
		pair := c.Pairs[matches[0].Index]
		c.Sequence.AddSequence(pair.A, pair.B)
		c.RemovePair(matches[0].Index)
		return c.Decode()
	}

	return getSolution(c)
}

func matchesShouldBeenSorted(matches []Match) bool {
	return len(matches) > 1
}

func getSolution(currentCase *Case) string {
	if currentCase.Sequence.AreEquals() {
		return NewPossibleSolution(currentCase.Sequence)
	}
	return NewImpossibleSolution()
}

func parseInputToCases(inputData []string) (cases []Case) {
	for _, data := range inputData {
		_, err := strconv.ParseInt(data, 10, 64)

		if err == nil {
			cases = append(cases, *NewCase())
			continue
		}

		splitData := strings.Split(data, " ")
		cases[len(cases)-1].AddPair(splitData[0], splitData[1])
	}
	return
}

type Case struct {
	Pairs    []Pair
	Sequence Sequence
	Solution string
}

type Pair struct {
	A, B string
}

type Sequence struct {
	A, B string
}

type Match struct {
	Index int
	Value string
}

func NewCase() *Case {
	return &Case{}
}

func NewPair(a, b string) *Pair {
	return &Pair{a, b}
}

func NewMatch(index int, value string) *Match {
	return &Match{index, value}
}

func NewImpossibleSolution() string {
	return impossibleSolution
}

func NewPossibleSolution(sequences Sequence) string {
	return sequences.A
}

func (c *Case) AddPair(a, b string) {
	c.Pairs = append(c.Pairs, *NewPair(a, b))
}

func (c *Case) RemovePair(index int) {
	var newPairs []Pair

	for i, pair := range c.Pairs {
		if i == index {
			continue
		}
		newPairs = append(newPairs, pair)
	}

	c.Pairs = newPairs
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

func (s *Sequence) AddSequence(a, b string) {
	s.A += a
	s.B += b
}

func (s *Sequence) AreEquals() bool {
	return len(s.A) > 0 && len(s.B) > 0 && strings.EqualFold(s.A, s.B)
}

func sortMatches(matches []Match) []Match {
	var (
		maxLength = int(^uint(0) >> 1)
		indexes   []int
	)

	for i, match := range matches {
		if len(match.Value) == maxLength {
			indexes = append(indexes, i)
		} else if len(match.Value) < maxLength {
			maxLength = len(match.Value)
			indexes = []int{i}
		}
	}

	if len(indexes) > 1 {
		var reducedMatches []Match
		for _, index := range indexes {
			reducedMatches = append(reducedMatches, matches[index])
		}
		matches = sortMatchesLexicographically(reducedMatches)
	}

	return matches
}

func sortMatchesLexicographically(matches []Match) []Match {
	for i := 0; i < len(matches)-1; i++ {
		for j := 0; j < len(matches)-i-1; j++ {
			if matches[j].Value > matches[j+1].Value {
				matches[j], matches[j+1] = matches[j+1], matches[j]
			}
		}
	}
	return matches
}

func readDataFromKeyboard() (inputData []string) {
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
