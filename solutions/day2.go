package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func DayTwo() *cobra.Command {
	var inputFile string
	cmd := &cobra.Command{
		Use:   "day2",
		Short: "Day 2 Solution",
		RunE: func(_ *cobra.Command, _ []string) error {
			return dayTwo(inputFile)
		},
	}
	cmd.Flags().StringVar(&inputFile, "input-file", "", "input file")
	return cmd
}

func dayTwo(inputFile string) error {
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	type Choice uint8
	const (
		Unknown  Choice = 0
		Rock     Choice = 1
		Paper    Choice = 2
		Scissors Choice = 3
	)

	toChoice := func(s string) Choice {
		switch s {
		case "A", "X":
			return Rock
		case "B", "Y":
			return Paper
		case "C", "Z":
			return Scissors
		}
		return Unknown
	}

	response := func(opponent Choice, result string) Choice {
		switch result {
		// need to lose
		case "X":
			switch opponent {
			case Rock:
				return Scissors
			case Paper:
				return Rock
			case Scissors:
				return Paper
			}
		// need to draw
		case "Y":
			return opponent
		// need to win
		case "Z":
			switch opponent {
			case Rock:
				return Paper
			case Paper:
				return Scissors
			case Scissors:
				return Rock
			}
		}
		return Unknown
	}

	score := func(opponent, response Choice) int {
		score := 0
		switch response {
		case Rock:
			score += 1
		case Paper:
			score += 2
		case Scissors:
			score += 3
		}
		if opponent == response {
			score += 3
		} else if opponent == Rock && response == Paper {
			score += 6
		} else if opponent == Paper && response == Scissors {
			score += 6
		} else if opponent == Scissors && response == Rock {
			score += 6
		}
		return score
	}

	finalScorePartOne := 0
	finalScorePartTwo := 0
	for scanner.Scan() {
		game := strings.Fields(scanner.Text())
		finalScorePartOne += score(toChoice(game[0]), toChoice(game[1]))
		finalScorePartTwo += score(toChoice(game[0]), response(toChoice(game[0]), game[1]))
	}

	fmt.Printf("Part One: %d\n", finalScorePartOne)
	fmt.Printf("Part Two: %d\n", finalScorePartTwo)

	return nil
}
