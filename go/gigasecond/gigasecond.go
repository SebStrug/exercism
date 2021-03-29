// Deal with gigaseconds
package gigasecond

// import path for the time package from the standard library
import "time"

// Add a gigasecond to any time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1e9)
	return t
}
