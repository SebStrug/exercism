package say

import (
	"fmt"
	"strings"
)

var wordMap = map[int64]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// spellTens spells the tens part of numbers from 0 - 1e2-1
func spellTens(num int64) string {
	if num == 0 {
		return ""
	}
	if spelling, ok := wordMap[num]; ok {
		return spelling
	}
	singleDigit := num % 10
	singleDigitSpelling := wordMap[singleDigit]
	tensDigitSpelling := wordMap[num-singleDigit]
	spelling := tensDigitSpelling + "-" + singleDigitSpelling
	return spelling
}

// spellHundreds spells the hundreds part of numbers from 0 - 1e3-1
func spellHundreds(numHundreds int64, tensDigit int64) string {
	spelling := ""
	if numHundreds > 0 {
		spelling = wordMap[numHundreds] + " hundred"
		if tensDigit > 0 {
			spelling = spelling + " " + spellTens(tensDigit)
		}
	}
	return spelling
}

// spellDigits spells the base (billions, millions, thousands) part
// of a number from 0 - 1e12-1
func spellDigits(numDigits int64, base string) string {
	if numDigits == 0 {
		return ""
	}

	spelling := ""
	if numDigits > 99 {
		numHundreds := (numDigits % 1000) / 100
		tensDigit := numDigits % 100
		spelling = spellHundreds(numHundreds, tensDigit)
	} else if numDigits > 0 {
		spelling = wordMap[numDigits]
	}
	spelling += fmt.Sprintf(" %s", base)
	return spelling
}

// Say spells out a number from 0 - 1e15-1 in plain English
func Say(num int64) (string, bool) {
	if num < 0 || num > 1000000000000-1 {
		return "", false
	}

	if num == 0 {
		return "zero", true
	}

	tensDigit := num % 100
	numHundreds := (num % 1000) / 100
	numThousands := (num % 1000000) / 1000
	numMillions := (num % 1000000000) / 1000000
	numBillions := num / 1000000000

	spelling := ""
	spelling += spellDigits(numBillions, "billion")
	if numMillions > 0 {
		spelling += " "
		spelling += spellDigits(numMillions, "million")
	}
	if numThousands > 0 {
		spelling += " "
		spelling += spellDigits(numThousands, "thousand")
	}
	if numHundreds > 0 {
		spelling += " "
		spelling += spellDigits(numHundreds, "hundred")
	}
	if tensDigit > 0 {
		spelling += " "
		spelling += spellTens(tensDigit)
	}
	return strings.TrimSpace(spelling), true
}
