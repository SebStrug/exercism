package protein

import (
	"errors"
)

var ErrStop = errors.New("Terminating codon")
var ErrInvalidBase = errors.New("No such codon")

// FromCodon converts a codon to its protein name
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

}

// Chunks splits a string into chunks of length chunkSize
// https://stackoverflow.com/a/61469854
func Chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

// FromRNA splits a string into its codons and returns the
// corresponding proteins
func FromRNA(RNA string) ([]string, error) {
	chunkedRNA := Chunks(RNA, 3)
	proteins := []string{}
	for _, chunk := range chunkedRNA {
		protein, err := FromCodon(chunk)
		if err == ErrStop {
			break
		} else if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
