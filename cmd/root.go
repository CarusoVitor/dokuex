package cmd

import (
	"fmt"
	"os"

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
		nameToFlag := make(map[string]string, cmd.Flags().NFlag())
		cmd.Flags().Visit(func(f *pflag.Flag) {
			nameToFlag[f.Name] = cmd.Flag(f.Name).Value.String()
		})
		pokemons, err := characteristics.MatchEmAll(nameToFlag)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
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
	// TODO: add more characteristics
	matchCmd.Flags().String("type", "", "Type of the pokemon")
	matchCmd.Flags().String("generation", "", "Generation of the pokemon in roman numerals")
	rootCmd.AddCommand(matchCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
