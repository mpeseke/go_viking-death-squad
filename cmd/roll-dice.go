package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var rollDiceCmd = &cobra.Command{
	Use:     "roll",
	Aliases: []string{"roll dice"},
	Short:   "Roll any amount of many-sided dice",
	Long:    "Rolls a user specified amount of a user-selected die",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		// Convert args from strings to integers
		lowestValue, err1 := strconv.Atoi(args[0])
		sides, err2 := strconv.Atoi(args[1])
		numberOfDice, err3 := strconv.Atoi(args[2])

		// Handle potential conversion errors
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Error: all arguments must be integers.")
			return
		}

		// Call the roll_dice function with integer arguments
		diceResults, total := roll_dice(lowestValue, sides, numberOfDice)

		// Print the results
		fmt.Printf("Rolling %d %d-sided dice, with lowest value %d\n", numberOfDice, sides, lowestValue)
		fmt.Printf("Dice Results: %v\n", diceResults)
		fmt.Printf("Total: %d\n", total)
	},
}

func init() {
	rootCmd.AddCommand(rollDiceCmd)
}
