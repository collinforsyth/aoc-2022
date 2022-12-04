package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func DayFour() *cobra.Command {
	var inputFile string
	cmd := &cobra.Command{
		Use:   "day4",
		Short: "Day 4 Solution",
		RunE: func(_ *cobra.Command, _ []string) error {
			return dayFour(inputFile)
		},
	}
	cmd.Flags().StringVar(&inputFile, "input-file", "", "input file")
	return cmd
}

func dayFour(inputFile string) error {
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	type interval struct{ begin, end int }
	parseInterval := func(i string) (interval, error) {
		s := strings.Split(i, "-")
		if len(s) != 2 {
			return interval{}, fmt.Errorf("invalid interval fmt: %s", s)
		}
		begin, err := strconv.Atoi(s[0])
		if err != nil {
			return interval{}, err
		}
		end, err := strconv.Atoi(s[1])
		if err != nil {
			return interval{}, err
		}
		return interval{begin: begin, end: end}, nil
	}

	fullyContains := func(x, y interval) bool {
		// with two intervals, x fully contains y iff the start of x is lower
		// than the start of y inclusive, and the end of x is greater than y inclusive.
		return x.begin <= y.begin && x.end >= y.end
	}

	overlaps := func(x, y interval) bool {
		return x.begin <= y.end && y.begin <= x.end
	}

	fullyContainsCounter := 0
	overlapCounter := 0
	for scanner.Scan() {
		partners := strings.Split(scanner.Text(), ",")
		pOne, err := parseInterval(partners[0])
		if err != nil {
			return err
		}
		pTwo, err := parseInterval(partners[1])
		if err != nil {
			return err
		}
		if fullyContains(pOne, pTwo) || fullyContains(pTwo, pOne) {
			fullyContainsCounter++
		}
		if overlaps(pOne, pTwo) {
			overlapCounter++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("Part One: %d\n", fullyContainsCounter)
	fmt.Printf("Part Two: %d\n", overlapCounter)

	return nil
}
