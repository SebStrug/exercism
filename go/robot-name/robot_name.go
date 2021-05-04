package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const maxNames int = 676000

// Robot is a real life machine
type Robot struct {
	name string
}

var existingNames = make(map[string]bool)

// init initialises a random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate a string of the form AA999
func generateName() string {
	nums := fmt.Sprintf("%03d", rand.Intn(1000))
	letters1 := rand.Intn(26) + 'A'
	letters2 := rand.Intn(26) + 'A'
	return fmt.Sprintf("%c%c%v", letters1, letters2, nums)
}

// Name generates a name for a Robot object if it doesn't have one
func (r *Robot) Name() (string, error) {
	if len(existingNames) == maxNames {
		return "", errors.New("exhausted all possible names")
	}
	if len(r.name) > 0 {
		return r.name, nil
	}

	r.name = generateName()
	for existingNames[r.name] {
		r.name = generateName()
	}
	existingNames[r.name] = true

	return r.name, nil
}

// Reset returns a Robot object with no name
func (r *Robot) Reset() {
	r.name = ""
}
