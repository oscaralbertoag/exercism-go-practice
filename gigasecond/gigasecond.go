package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	const Gigasecond int32 = 1000000000
	return t.Add(time.Duration(Gigasecond) * time.Second)
}
