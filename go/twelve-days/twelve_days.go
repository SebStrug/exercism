package twelve

import (
	"fmt"
	"strings"
)

var days = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}
var gifts = [12]string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"and a Partridge in a Pear Tree.",
}

// Verse outputs a line from 'The Twelve Days of Christmas'
// indexed from 1 not 0
func Verse(num int) string {
	s := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[num-1])

	for i := 12 - num; i < 12; i++ {
		if num == 1 {
			s += "a Partridge in a Pear Tree."
		} else {
			s += gifts[i]
			s += ", "
		}
	}
	return strings.TrimSuffix(s, ", ")
}

// Song outputs the full lyrics to 'The Twelve Days of Christmas'
func Song() string {
	song := ""
	for i := 0; i < 12; i++ {
		song += fmt.Sprintf("%s\n", Verse(i+1))
	}
	return strings.TrimSuffix(song, "\n")
}
