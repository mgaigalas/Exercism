// Package gigasecond contains tools to add one Giga second for provided time
package gigasecond

import "time"

const GigaSecond time.Duration = 1e9 * time.Second

// AddGigasecond adds Giga second to provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(GigaSecond)
}
