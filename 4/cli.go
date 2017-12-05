package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//CLI is stream.
type CLI struct {
	outStream, errStream io.Writer
}

//Run is function args parse and diff.
func (c *CLI) Run(args []string) int {
	files := []string{}
	var u bool
	flags := flag.NewFlagSet("mdiff", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&u, "u", false, "Output change + and -.")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}

	for 0 < flags.NArg() {
		files = append(files, flags.Args()[0])

		err := flags.Parse(flags.Args()[1:])
		if err != nil {
			fmt.Sprintln(c.errStream, err)
			return 1
		}
	}

	af, err := os.Open(files[0])
	defer af.Close()
	a, err := ioutil.ReadAll(af)

	bf, err := os.Open(files[1])
	defer bf.Close()
	b, err := ioutil.ReadAll(bf)

	if err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}

	fmt.Fprint(c.outStream, diff(string(a), string(b), u))

	return 0
}
