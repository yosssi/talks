package main

func main() {
	// START OMIT
	var c <-chan bool
	close(c) //-> invalid operation: close(c) (cannot close receive-only channel) // HL
	// END OMIT
}
