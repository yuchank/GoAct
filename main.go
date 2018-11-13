package main

import (
	"log"
	"os"

	_ "goact/matchers"
	"goact/search"
)

// init is called prior to main.
func init() {
	// change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

// As of Go 1.11,
// $ go mod init <mod_name>
// $ import "mod_name/pkg_name"
