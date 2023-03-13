package main
import "C"
import (
	"github.com/moond4rk/HackBrowserData/cmd/hack-browser-data"
)

//export run
func run() {
	cmd.Execute()
}

