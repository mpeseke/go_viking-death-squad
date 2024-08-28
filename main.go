/*
Copyright © 2024 Marlon Peseke mpeseke@gmail.com
*/
package main

import (
	"bufio"
	"fmt"
	"go_viking-death-squad/cmd"
	"os"
	"strings"
)

type Character struct {
	Race         string
	Name         string
	SpecialSkill string
	Oath         string
	Debt         int
	Code         string
	Power        int
	Aim          int
	Speed        int
	Wits         int
	Guts         int
	Resolve      int
	Skills       []string
	Gear         []string
	Armor        []string
	BloodRunes   []string
	Desc         string
}

func main() {

	cmd.Execute()
	fmt.Println("Welcome to Viking Death Squad")
	choose_race()
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
