package main

// START OMIT
// After waits for the duration to elapse and then sends the current time
// on the returned channel.
// It is equivalent to NewTimer(d).C.
func After(d Duration) <-chan Time { // HL
	return NewTimer(d).C
}

// NewTimer creates a new Timer that will send
// the current time on its channel after at least duration d.
func NewTimer(d Duration) *Timer {
	c := make(chan Time, 1) // HL
	t := &Timer{
		C: c,
		// ...
	}
	startTimer(&t.r)
	return t
}

// END OMIT
