package main

import (
// "log"
// "os"
)

const VERSION string = "0.0.1"

const USAGE string = `go-swg-gen is a CLI tool for generating Swagger 2.0 documentation from Commands.

Usage:
  go-swg-gen [options] <command> <command>...

Options:
  -h, --help    Show this screen.
  -v, --version Show version.
`

func main() {
	print(USAGE)
}
