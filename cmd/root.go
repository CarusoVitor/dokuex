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
		characteristicsMap := make(map[string]characteristics.Characteristic, cmd.Flags().NFlag())
		cmd.Flags().Visit(func(f *pflag.Flag) {
			characteristicsMap[f.Name] = characteristics.NewCharacteristic(f.Name)
		})
		pokemons := make(map[string]struct{}, 0)
		for flag, characteristic := range characteristicsMap {
			value := cmd.Flag(flag).Value.String()
			result, err := characteristic.Get(value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting characteristic %s: %v\n. The match will proceed ignoring it", flag, err)
			}
			if len(result) == 0 {
				fmt.Printf("No pokemons found for characteristic %s with value %s\n", flag, value)
				break
			}
			if len(pokemons) == 0 {
				pokemons = make(map[string]struct{}, len(result))
				for name := range result {
					pokemons[name] = struct{}{}
				}
			} else {
				pokemons = intersect(pokemons, result)
				if len(pokemons) == 0 {
					fmt.Println("No pokemons found matching all characteristics")
					break
				}
			}
		}
	},
}

// TODO: implement set intersection
func intersect(a, b map[string]struct{}) map[string]struct{} {
	return a
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
