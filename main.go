package main

import (
	"log"
	"os"

	_ "matchers"
	"search"
)

// init is called prior to main.
func init() {
	// change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
