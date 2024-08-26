package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

func parseFlags() string {
	const usage = `Usage of hungarian-lottery %s:

    hungarian-lottery <file>

Options:
    - file: file to read (required)
`
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, getVersion())
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	validateFlags()
	return os.Args[1]
}

func validateFlags() {
	if len(os.Args) < 2 {
		fmt.Println("File is required")
		flag.Usage()
	}
}

func getVersion() string {
	if i, ok := debug.ReadBuildInfo(); ok {
		return i.Main.Version
	}
	return "(unknown)"
}
