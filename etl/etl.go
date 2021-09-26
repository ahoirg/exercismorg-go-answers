package etl

import "strings"

func Transform(input map[int][]string) (score map[string]int) {
	score = make(map[string]int)

	for _, items := range input {
		for _, item := range items {
			item = strings.ToLower(item)
			switch item {
			case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
				score[item] = 1
			case "d", "g":
				score[item] = 2
			case "b", "c", "m", "p":
				score[item] = 3
			case "f", "h", "v", "w", "y":
				score[item] = 4
			case "k":
				score[item] = 5
			case "j", "x":
				score[item] = 8
			case "q", "z":
				score[item] = 10
			}
		}
	}

	return
}
