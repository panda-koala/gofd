package main

import "flag"

type Arg struct {
	url *string
	dir *string
	chuck *int
	args []string
}

func (a *Arg) initArg() {
	a.url = flag.String("u", "", "file http url")
	a.dir = flag.String("d", "", "download local path")
	a.chuck = flag.Int("c", 7, "download local path")
	a.args = flag.Args()
	flag.Parse()
}


