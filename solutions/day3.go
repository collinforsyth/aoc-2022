package solutions

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/spf13/cobra"
)

func DayThree() *cobra.Command {
	var inputFile string
	cmd := &cobra.Command{
		Use:   "day3",
		Short: "Day 3 Solution",
		RunE: func(_ *cobra.Command, _ []string) error {
			return dayThree(inputFile)
		},
	}
	cmd.Flags().StringVar(&inputFile, "input-file", "", "input file")
	return cmd
}

func dayThree(inputFile string) error {
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	compartmentSplit := func(s string) (string, string) {
		return s[:len(s)/2], s[len(s)/2:]
	}

	findCommonItem := func(input ...string) rune {
		// returns the first common item in both strings
		counter := make(map[rune]map[int]int)
		for i, s := range input {
			for _, r := range s {
				if _, ok := counter[r]; !ok {
					counter[r] = map[int]int{i: 1}
				} else {
					counter[r][i] += 1
				}
			}
		}
		commonCount := len(input)
		for k, v := range counter {
			if len(v) == commonCount {
				return k
			}
		}
		return 0
	}

	itemPriority := func(item rune) int {
		if item == rune(0) {
			return 0
		}
		if unicode.IsUpper(item) {
			return 52 - (90 - int(item))
		} else {
			return 26 - (122 - int(item))
		}
	}

	prioritySum := 0
	badgePriority := 0
	badgeCollector := make([]string, 0, 3)
	for scanner.Scan() {
		x, y := compartmentSplit(scanner.Text())
		common := findCommonItem(x, y)
		prioritySum += itemPriority(common)
		badgeCollector = append(badgeCollector, scanner.Text())
		if len(badgeCollector) == 3 {
			commonElf := findCommonItem(badgeCollector...)
			badgePriority += itemPriority(commonElf)
			badgeCollector = badgeCollector[:0]
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("Part One: %d\n", prioritySum)
	fmt.Printf("Part Two: %d\n", badgePriority)

	return nil
}
