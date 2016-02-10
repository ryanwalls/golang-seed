package main

import (
	"github.com/ryanwalls/golang-seed/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
