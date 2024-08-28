package cmd

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func roll_dice(min, max, no_of_die int) ([]int, int) {
	var rolls []int

	for i := 0; i < no_of_die; i++ {
		rolls = append(rolls, rand.Intn((max+1)-min)+min)
	}

	bool := check_for_crit(rolls)

	if bool {
		fmt.Printf("Six-six-six, the Number of the Beast \n Hel and fire was spawned to be released! \n")
	}

	total := 0

	for _, rolls := range rolls {
		total += rolls
	}

	return rolls, total
}

func check_for_crit(dice_roll []int) bool {
	if len(dice_roll) < 3 {
		return false
	}

	count := 0

	for _, value := range dice_roll {
		if value == 6 {
			count++
		}

		if count >= 3 {
			return true
		}
	}

	return false
}
