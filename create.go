package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var args struct {
	path string
	size string
}

func create(path, size string) error {
	dev, err := freeDev()
	if err != nil {
		return err
	}

	cmds := [][]string{
		[]string{"fallocate", "-l", size, path},
		[]string{"losetup", dev, path},
		[]string{"tcplay", "-c", "-d", dev, "-a", "whirlpool", "-b", "AES-256-XTS"},
	}
	fmt.Printf("%v\n", cmds)
	for _, args := range cmds {
		var stderr bytes.Buffer
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = &stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return fmt.Errorf(stderr.String())
		}
	}
	return nil
}

func init() {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	fs.StringVar(&args.path, "path", "", "path to crypt")
	fs.StringVar(&args.size, "size", "", "size of crypt (e.g. 3G)")

	AddCmd("create", fs, func() error { return create(args.path, args.size) })
}
