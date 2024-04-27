// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes in a time, and adds 10^9 seconds to that time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(1000000000) * time.Second)
	return t
}
