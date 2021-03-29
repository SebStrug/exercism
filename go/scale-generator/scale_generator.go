package scale

import "strings"

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"s", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
// Scale generates a musical scale given the starting note/tonic
// and the associated interval(s)
func Scale(tonic, interval string) []string {
	tonic = strings.ToUpper(tonic)
}