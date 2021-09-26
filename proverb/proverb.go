package proverb

import (
	"strings"
)

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	var proverb []string
	for i := 0; i < len(rhyme)-1; i++ {
		r := strings.NewReplacer(
			"$1", rhyme[i],
			"$2", rhyme[i+1],
		)
		proverb = append(proverb, r.Replace("For want of a $1 the $2 was lost."))
	}
	proverb = append(proverb, strings.Replace("And all for the want of a $1.", "$1", rhyme[0], 1))

	return proverb
}
