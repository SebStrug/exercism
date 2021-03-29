package clock

import "fmt"

// Clock takes hour, minute input
type Clock struct{ h, m int }

// New produces a new clock format
func New(h, m int) Clock {
	if m >= 0 {
		// handle overflow housr by simple padding
		h = (h + (100 * 24) + (m / 60)) % 24
		m = m % 60
		return Clock{h, m}
	}

	if m < -60 {
		h += (m / 60) - 1
	} else {
		h--
	}
	h = (h + (300 * 24)) % 24
	m = (m + (300 * 60)) % 60
	return Clock{h, m}
}

// String returns a readable clock time
func (t Clock) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}

// Add adds minutes to a clock
func (t Clock) Add(m int) Clock {
	return New(t.h, t.m+m)
}

// Subtract subtracts minutes from a clock
func (t Clock) Subtract(m int) Clock {
	return New(t.h, t.m-m)
}
