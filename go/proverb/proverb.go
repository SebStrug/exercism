package proverb

import (
	"fmt"
)

// Proverb generates a proverb given a list of inputs
func Proverb(rhyme []string) []string {
	rhymeLen := len(rhyme)
	if rhymeLen == 0 {
		return rhyme
	}

	prov := []string{}
	for ind := range rhyme {
		line := ""
		if ind == rhymeLen-1 {
			line = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		
		}else {
			line = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[ind], rhyme[ind+1])
		}
		prov = append(prov, line)
	}
	return prov
}
