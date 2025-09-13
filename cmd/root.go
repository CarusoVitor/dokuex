package cmd

import (
	"fmt"
	"os"

	"github.com/CarusoVitor/dokuex/characteristics"
	"github.com/CarusoVitor/dokuex/pokeapi"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "dokuex",
	Short: "A pokedoku solver",
	Long:  `dokuex is a pokedoku solver that queries pokeapi to find the answers of each characteristic combination`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

type cobraFlag struct {
	name  string
	value any
}

func processFlagsValues(cmd *cobra.Command) (map[string][]string, error) {
	nameToFlags := make(map[string][]string, cmd.Flags().NFlag())
	var err error = nil

	cmd.Flags().Visit(func(f *pflag.Flag) {
		var values []string

		switch f.Value.Type() {
		case "string":
			var value string
			value, err = cmd.Flags().GetString(f.Name)
			values = []string{value}
		case "stringSlice":
			values, err = cmd.Flags().GetStringSlice(f.Name)
		case "bool":
			var flag bool
			flag, err = cmd.Flags().GetBool(f.Name)
		}

		nameToFlags[f.Name] = values

	})
	return nameToFlags, err
}

// TODO: allow multiple runs to have an actual cache usage
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Match pokedoku characteristics",
	Long:  "A command that returns all pokemon that match the given characteristics",
	PreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Println("Please provide at least one characteristic flag")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		nameToFlags := processFlagsValues(cmd)

		client := pokeapi.NewPokeApiClient()
		pokemons, err := characteristics.MatchEmAll(nameToFlag, client)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if len(pokemons) == 0 {
			fmt.Println("No pokemons found matching the given characteristics")
			return
		}
		fmt.Println("Pokemons found:")
		i := 0
		for pokemon := range pokemons {
			i++
			fmt.Printf("%d. %s\n", i, pokemon)
		}
	},
}

func init() {
	matchCmd.Flags().StringSlice("type", []string{}, "Type of the pokemon")
	matchCmd.Flags().String("generation", "", "Generation of the pokemon in the form generation-Z, where Z is a roman numeral from 1 to 9")
	matchCmd.Flags().StringSlice("move", []string{}, "TM or HM in pokemon games")
	matchCmd.Flags().StringSlice("ability", []string{}, "Ability of the pokemon (including hidden)")
	matchCmd.Flags().Bool("ultra-beast", true, "Ultra beast pokemons")
	rootCmd.AddCommand(matchCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
