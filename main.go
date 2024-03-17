package main

import (
	"fmt"
	"os"

	"learn_projects/hotel"
	"learn_projects/sudoku"
	"learn_projects/sum"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "learn_projects",
	Short: "learn_projects",
}

var sudokuCmd = &cobra.Command{
	Use:   "sudoku",
	Short: "sudoku",
	Run: func(cmd *cobra.Command, args []string) {
		sudoku.Sudoku()
	},
}

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "sum",
	Run: func(cmd *cobra.Command, args []string) {
		sum.Sum()
	},
}

var hotelCmd = &cobra.Command{
	Use:   "hotel",
	Short: "hotel",
	Run: func(cmd *cobra.Command, args []string) {
		hotel.Hotel()
	},
}

func main() {
	rootCmd.AddCommand(sudokuCmd)
	rootCmd.AddCommand(sumCmd)
	rootCmd.AddCommand(hotelCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
