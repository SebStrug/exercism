package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot is a real life machine
type Robot struct {
	name string
}

const capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var existingNames = make(map[string]bool)

// init initialises a random seed
func init(){
	rand.Seed(time.Now().UnixNano())
}

// Generate a string of the form AA999
func generateName() string {
	nums := rand.Intn(899)+100
	letters := make([]byte, 2)
    for i := range letters {
        letters[i] = capitalLetters[rand.Intn(len(capitalLetters))]
    }
	return fmt.Sprintf("%s%d", letters, nums)
}

// Name generates a name for a Robot object if it doesn't have one
func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	for {
		newName := generateName()
		if _, ok := existingNames[newName]; !ok {
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
	return
}