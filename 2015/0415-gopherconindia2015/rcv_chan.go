package main

import (
	"os"
	"os/signal"
)

func main() {
	// START OMIT
	sigs := make(chan os.Signal, 10) // HL
	go handleSigs(sigs)
	signal.Notify(sigs, os.Interrupt) // HL
	// later ...
	signal.Notify(sigs, os.Kill) // HL
	// END OMIT
}
