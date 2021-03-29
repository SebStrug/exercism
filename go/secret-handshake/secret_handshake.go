package secret

import (
	"strconv"
)

var shakes = []string{"wink", "double blink", "close your eyes", "jump", ""}

func reverse(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}

// Handshake produces a secret string handshake from a code
func Handshake(n uint) []string {
	code := strconv.FormatInt(int64(n), 2)
	codeLen := len(code)
	if codeLen < 1 {
		return []string{}
	}

	var res []string
	for i := codeLen - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(code[i]))
		if val == 1 {
			res = append(res, shakes[codeLen-1-i])
		}
	}

	for ind, item := range res {
		if item == "" {
			fixedRes := append(res[:ind], res[ind+1:]...)
			fixedRes = reverse(fixedRes)
			return fixedRes
		}
	}
	return res
}
