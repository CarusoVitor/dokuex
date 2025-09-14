package cmd

import (
	"fmt"
	"os"

	"github.com/CarusoVitor/dokuex/api"
	"github.com/CarusoVitor/dokuex/characteristics"
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

type characteristicFlag struct {
	name  string
	value any
}

// processFlagsValues formats active key-value flag pairs to a map
func processFlagsValues(cmd *cobra.Command) (map[string][]string, error) {
	nameToFlags := make(map[string][]string, cmd.Flags().NFlag())
	var err error = nil

	cmd.Flags().Visit(func(f *pflag.Flag) {
		var values []string

		flagType := f.Value.Type()

		switch flagType {
		case "string":
			var value string
			value, err = cmd.Flags().GetString(f.Name)
			values = []string{value}
		case "stringSlice":
			values, err = cmd.Flags().GetStringSlice(f.Name)
		case "bool":
			values = []string{"true"}
		default:
			values, err = nil, fmt.Errorf("flag type %s wasn't yet implemented", flagType)
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
		nameToValues, err := processFlagsValues(cmd)
		if err != nil {
			panic(err)
		}

		client := api.NewPokeApiClient()
		pokemons, err := characteristics.MatchEmAll(nameToValues, client)

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
	matchCmd.Flags().String("generation", "", "Pokemon generation in generation-Z form, where Z is a roman numeral from 1 to 9")
	matchCmd.Flags().StringSlice("move", []string{}, "Pokemon moves")
	matchCmd.Flags().StringSlice("ability", []string{}, "Pokemon abilities (including hidden)")
	matchCmd.Flags().Bool("ultra-beast", true, "Ultra beast pokemons")
	rootCmd.AddCommand(matchCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
