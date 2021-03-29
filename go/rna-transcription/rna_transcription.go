package strand

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts a DNA strand to its transcribed RNA strand
func ToRNA(dna string) string {
	RNA := make([]rune, len(dna))
	for ind, val := range dna {
		RNA[ind] = complements[val]
	}
	return string(RNA)
}
