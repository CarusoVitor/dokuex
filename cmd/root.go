package cmd

import (
	"fmt"
	"os"

	"github.com/CarusoVitor/dokuex/characteristics"
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
	matchCmd.Flags().StringSlice(characteristics.TypeName, []string{}, "Type of the pokemon")
	matchCmd.Flags().String(characteristics.GenerationName, "", "Pokemon generation in generation-Z form, where Z is a roman numeral from 1 to 9")
	matchCmd.Flags().StringSlice(characteristics.MoveName, []string{}, "Pokemon moves")
	matchCmd.Flags().StringSlice(characteristics.AbilityName, []string{}, "Pokemon abilities (including hidden)")
	matchCmd.Flags().Bool(characteristics.UltraBeastName, true, "Ultra beast pokemons")
	matchCmd.Flags().Bool(characteristics.MegaName, true, "Mega pokemons and their base forms")
	matchCmd.Flags().Bool(characteristics.GmaxName, true, "Gigantamax pokemons and their base forms")
	matchCmd.Flags().Bool(characteristics.LegendaryName, true, "Legendary pokemons")
	matchCmd.Flags().Bool(characteristics.BabyName, true, "Baby pokemons")
	matchCmd.Flags().Bool(characteristics.MythicalName, true, "Mythical pokemons")
	rootCmd.AddCommand(matchCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
