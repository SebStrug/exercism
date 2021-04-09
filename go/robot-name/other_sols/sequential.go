package robotname

// This second iteration doesn't return a random name per se,
//the names are sequential and predictable. However it does use
//the whole namespace. Returning a random name each time would be
//much slower as I would have to create and check a map constantly,
//generating random names until I generated a name not in the map.

import (
	"errors"
	"fmt"
)

type Robot struct {
	name string
}

var capitalLetters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var firstLetterIndex = 0
var secondLetterIndex = 0
var robotNumber = 0
var secondRound = false

// Generate a string of the form AA999
func generateName() (string, error) {
	paddedNum := fmt.Sprintf("%03d", robotNumber)
	name := fmt.Sprintf("%s%s%s", capitalLetters[firstLetterIndex], capitalLetters[secondLetterIndex], paddedNum)
	if robotNumber > 999 {
		robotNumber = 0
		firstLetterIndex++
	}
	if firstLetterIndex >= 26 {
		firstLetterIndex = 0
		secondLetterIndex++
	}
	if secondLetterIndex >= 26 {
		secondLetterIndex = 0
		secondRound = true
	}
	robotNumber++

	if secondRound && firstLetterIndex == 0 && secondLetterIndex == 0 && robotNumber == 1 {
		return "", errors.New("Namespace exhausted")
	}
	return name, nil
}

func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	name, err := generateName()
	r.name = name
	return r.name, err
}

func (r *Robot) Reset() {
	r.name = ""
	return
}
