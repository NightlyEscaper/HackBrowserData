package main
import "C"
import (
	"cmd/hack-browser-data"
)

//export run
func run() {
	cmd.Execute()
}

