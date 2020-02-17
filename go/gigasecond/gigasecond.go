//Package gigasecond adds a gigasecond to current time
package gigasecond

import (
	"time"
)

// AddGigasecond to provided date
func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * 1000000000)
}
