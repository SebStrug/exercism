package strain

type Ints []int
type Strings []string
type Lists [][]int
		
// Keep returns the subset of a collection that fulfils a predicate
func (coll Ints) Keep(pred func(int) bool) Ints {
	var keep Ints
	for _, val := range coll {
		if pred(val) {
			keep = append(keep, val)
		}
	}
	return keep
}

// Discard returns the subset of a collection that does not fulfil a predicate
func (coll Ints) Discard(pred func(int) bool) Ints {
	var keep Ints
	for _, val := range coll {
		if !pred(val) {
			keep = append(keep, val)
		}
	}
	return keep
}

func (coll Strings) Keep(pred func(string) bool) Strings {
	var keep Strings
	for _, val := range coll {
		if pred(val) {
			keep = append(keep, val)
		}
	}
	return keep
}

func (coll Lists) Keep(pred func([]int) bool) Lists {
	var keep Lists
	for _, val := range coll {
		if pred(val) {
			keep = append(keep, val)
		}
	}
	return keep
}