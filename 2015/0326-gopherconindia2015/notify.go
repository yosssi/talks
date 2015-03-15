package main

import "os"

// START OMIT
// Notify causes package signal to relay incoming signals to c.
// ...
// Package signal will not block sending to c ... // HL
// ...
func Notify(c chan<- os.Signal, sig ...os.Signal) {
	// END OMIT
}
