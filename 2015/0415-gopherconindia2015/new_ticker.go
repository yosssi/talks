package main

// START OMIT
// NewTicker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
// It adjusts the intervals or drops ticks to make up for slow receivers. // HL
// ...
func NewTicker(d Duration) *Ticker {
	// END OMIT
}
