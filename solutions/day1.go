package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

func DayOne() *cobra.Command {
	var inputFile string
	cmd := &cobra.Command{
		Use:   "day1",
		Short: "Day One Solution",
		RunE: func(_ *cobra.Command, _ []string) error {
			return dayOne(inputFile)
		},
	}
	cmd.Flags().StringVar(&inputFile, "input-file", "", "input file")
	return cmd
}

func dayOne(inputFile string) error {
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	maxCalories := math.MinInt64
	currentCalories := 0

	// for part two, naive approach is to put all totals in an array
	// and then sort
	var calorieCount = []int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			calorieCount = append(calorieCount, currentCalories)
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}
		cals, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		currentCalories += cals
	}

	sort.Ints(calorieCount)
	partTwoCount := 0
	for i := 0; i < 3; i++ {
		// fascinated if the effort to use the off the shelf integer
		// sorting and then counting backwards was less work than just
		// writing a decreasing sort function. probably not.
		partTwoCount += calorieCount[len(calorieCount)-i-1]
	}

	fmt.Printf("Part One: %d\n", maxCalories)
	fmt.Printf("Part Two: %d\n", partTwoCount)

	return nil
}
