package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func DayFive() *cobra.Command {
	var inputFile string
	cmd := &cobra.Command{
		Use:   "day5",
		Short: "Day 5 Solution",
		RunE: func(_ *cobra.Command, _ []string) error {
			return dayFive(inputFile)
		},
	}
	cmd.Flags().StringVar(&inputFile, "input-file", "", "input file")
	return cmd
}

type stack struct {
	s []rune
}

func (s *stack) Pop() rune {
	x := s.s[0]
	s.s = s.s[1:]
	return x
}

func (s *stack) Push(r rune) {
	s.s = append([]rune{r}, s.s...)
}

func (s *stack) Peek() rune {
	return s.s[0]
}

func dayFive(inputFile string) error {
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	intializeStacks := func(s *bufio.Scanner) ([]stack, error) {
		stackStrings := make([]string, 0)
		for s.Scan() {
			// on empty line we're done
			if s.Text() == "" {
				break
			}
			stackStrings = append(stackStrings, s.Text())
		}
		// initialStacks are last line, numbered 1 - 9
		numStackString := stackStrings[len(stackStrings)-1]
		numStacks := len(strings.Fields(numStackString))
		stacks := make([]stack, numStacks)
		for i := len(stackStrings) - 2; i >= 0; i-- {
			counter := 0
			for counter < len(stackStrings[i]) {
				// to parse, empty is 3 spaces, value is [X], where X is a rune
				// that goes on the stack
				val := stackStrings[i][counter : counter+3]
				if val != "   " {
					stacks[counter/4].Push(rune(val[1]))
				}
				counter += 4
			}
		}
		return stacks, nil
	}

	cloneStacks := func(s []stack) []stack {
		c := make([]stack, len(s))
		for i := range s {
			x := make([]rune, len(s[i].s))
			copy(x, s[i].s)
			c[i].s = x
		}
		return c
	}

	partOneStacks, err := intializeStacks(scanner)
	if err != nil {
		return err
	}
	partTwoStacks := cloneStacks(partOneStacks)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		num := MustAtoi(fields[1])
		src := MustAtoi(fields[3])
		dst := MustAtoi(fields[5])
		accum := make([]rune, num)
		for i := 0; i < num; i++ {
			pOne, pTwo := partOneStacks[src-1].Pop(), partTwoStacks[src-1].Pop()
			accum[i] = pTwo
			partOneStacks[dst-1].Push(pOne)
		}
		for i := num - 1; i >= 0; i-- {
			partTwoStacks[dst-1].Push(accum[i])
		}
	}
	var dayOneSolution = make([]rune, len(partOneStacks))
	for i := range partOneStacks {
		dayOneSolution[i] = partOneStacks[i].Peek()
	}
	var dayTwoSolution = make([]rune, len(partTwoStacks))
	for i := range partTwoStacks {
		dayTwoSolution[i] = partTwoStacks[i].Peek()
	}

	fmt.Printf("Part One: %s\n", string(dayOneSolution))
	fmt.Printf("Part Two: %s\n", string(dayTwoSolution))

	return nil
}
