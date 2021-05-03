package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Robot is a real life machine
type Robot struct {
	name string
}

var existingNames = make(map[string]bool)

var namePattern = regexp.MustCompile("[A-Z]{2}[0-9]{3}")

// init initialises a random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate a string of the form AA999
func generateName() string {
	nums := rand.Intn(899) + 100
	letters1 := rand.Intn(26) + 'A'
	letters2 := rand.Intn(26) + 'A'
	return fmt.Sprintf("%c%c%d", letters1, letters2, nums)
}

// Name generates a name for a Robot object if it doesn't have one
func (r *Robot) Name() (string, error) {
	if len(existingNames) == 676000 {
		return "", errors.New("exhausted all possible names")
	}
	if len(r.name) > 0 {
		if !namePattern.MatchString(r.name) {
			return "", errors.New("robot name is invalid")
		}
		return r.name, nil
	}
	for {
		newName := generateName()
		if !existingNames[newName] {
			existingNames[newName] = true
			r.name = newName
			break
		}
	}

	return r.name, nil
}

// Reset returns a Robot object with no name
func (r *Robot) Reset() {
	r.name = ""
}
