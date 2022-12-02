package main

import (
	"log"

	"github.com/collinforsyth/aoc-2022/solutions"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "aoc2022",
		Short: "Advent of Code 2022 Solutions",
	}
	cmd.AddCommand(solutions.DayOne())
	cmd.AddCommand(solutions.DayTwo())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
