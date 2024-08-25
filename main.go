package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, Go!")
	choose_race()

	fmt.Println(roll_dice(1, 6, 3))
}

func roll_dice(min, max, no_of_die int) ([]int, int) {
	var rolls []int

	for i := 0; i < no_of_die; i++ {
		rolls = append(rolls, rand.Intn((max+1)-min)+min)
	}

	total := 0

	for _, rolls := range rolls {
		total += rolls
	}

	return rolls, total
}

func choose_race() string {
	fmt.Println("Will you play as human, or Immortal? ")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please choose an option (type '1' or '2'):")
		fmt.Println("1: Human")
		fmt.Println("2: Immortal")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "1" || input == "2" {
			if input == "1" {
				fmt.Println("You have chosen to fight as the noble human, raise your fist and fight for death and glory in Valholl.")
				return "Human"
			}

			if input == "2" {
				fmt.Println("You have chosen to fight as the mighty Immortal, cursed to wander Urth for all eternity, swept to and fro by the whims of the gods.")
				return "Immortal"
			}
		}

		fmt.Println("Invalid choice. Please choose again. ")
	}

}
