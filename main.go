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
	fmt.Println(roll_dice(1, 20)) //simulate rolling d20
	fmt.Println(roll_dice(1, 6))  //simulate rolling d6
	choose_race()
}

func roll_dice(min, max int) int {
	return rand.Intn((max+1)-min) + min
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
			return input
		}

		fmt.Println("Invalid choice. Please choose again. ")
	}

}
