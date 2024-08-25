package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println(roll_dice(1, 20)) //simulate rolling d20
	fmt.Println(roll_dice(1, 6))  //simulate rolling d6
}

func roll_dice(min, max int) int {
	return rand.Intn((max+1)-min) + min

}
