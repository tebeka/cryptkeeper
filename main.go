// cryptkeeper manages secure volumes (crypts)

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func main() {
	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Printf("usage: %s <command> [<args>]\n", name)
		fmt.Printf("Manage secure volumes (crypts)\n")
		fmt.Printf("\nAvailable sub commands:\n")
		for _, sub := range SubCmds() {
			fmt.Printf("\t%s\n", sub)
		}
		flag.PrintDefaults()
	}

	name := os.Args[1]
	switch name {
	case "-h":
		flag.Usage()
		os.Exit(0)
	case "--help":
		flag.Usage()
		os.Exit(0)
	}

	cmd, ok := FindCmd(name)
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown command %s\n", name)
		os.Exit(1)
	}

	cmd.Parser.Parse(os.Args[2:]) // Might exit
	if err := cmd.Handler(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", name, err)
		os.Exit(1)
	}
}
