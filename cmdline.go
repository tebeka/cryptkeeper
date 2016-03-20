package main

import (
	"flag"
	"sort"
)

type Cmd struct {
	Parser  *flag.FlagSet
	Handler func() error
}

var cmds = make(map[string]*Cmd)

func AddCmd(name string, parser *flag.FlagSet, handler func() error) {
	cmds[name] = &Cmd{
		Parser:  parser,
		Handler: handler,
	}
}

func FindCmd(name string) (*Cmd, bool) {
	cmd, ok := cmds[name]
	return cmd, ok
}

func SubCmds() []string {
	var subs []string

	for name := range cmds {
		subs = append(subs, name)
	}
	sort.Strings(subs)
	return subs
}
