package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4
const gigaSecond = 1e9

func AddGigasecond(birth time.Time) time.Time {
	return time.Unix(birth.Unix()+gigaSecond, 0)
}
