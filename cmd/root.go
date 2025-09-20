package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dokuex",
	Short: "A pokedoku solver",
	Long:  `dokuex is a pokedoku solver that queries pokeapi to find the answers of each characteristic combination`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	matchCmd.Flags().StringSlice("type", []string{}, "Type of the pokemon")
	matchCmd.Flags().String("generation", "", "Pokemon generation in generation-Z form, where Z is a roman numeral from 1 to 9")
	matchCmd.Flags().StringSlice("move", []string{}, "Pokemon moves")
	matchCmd.Flags().StringSlice("ability", []string{}, "Pokemon abilities (including hidden)")
	matchCmd.Flags().Bool("ultra-beast", true, "Ultra beast pokemons")
	matchCmd.Flags().Bool("mega", true, "Mega pokemons")
	rootCmd.AddCommand(matchCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
