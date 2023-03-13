package main
import "C"
import (
	"github.com/moond4rk/HackBrowserData/cmd"
)

//export run
func run() {
	cmd.Execute()
}

