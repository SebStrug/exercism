package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// CountLetters sends all runes in a set of texts to a channel
func CountLetters(items []string, c chan rune) {
	for _, s := range items {
		for _, r := range s {
			c <- r
		}
	}
}

// ConcurrentFrequency counts the frequency of each rune in a text
// concurrently
func ConcurrentFrequency(items []string) FreqMap {
	f := FreqMap{}
	c := make(chan rune)
	defer close(c)
	go CountLetters(items, c)
	for i := range c {
		f[i]++
	}
	return f
}