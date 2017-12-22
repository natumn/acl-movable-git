package main

import (
<<<<<<< HEAD
	"flag"
=======
>>>>>>> upstream/master
	"fmt"
	"io"
	"io/ioutil"
	"os"
<<<<<<< HEAD
)

//CLI is stream.
=======
	"flag"
)

>>>>>>> upstream/master
type CLI struct {
	outStream, errStream io.Writer
}

<<<<<<< HEAD
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
=======
func (c *CLI) Run(args []string) int {
	filenames := []string{}

	var unified bool
	flags := flag.NewFlagSet("mdiff", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&unified, "u", false, "output as unified context")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}
	for 0 < flags.NArg() {
		filenames = append(filenames, flags.Args()[0])

		err := flags.Parse(flags.Args()[1:])
		if err != nil {
			fmt.Fprintln(c.errStream, err)
>>>>>>> upstream/master
			return 1
		}
	}

<<<<<<< HEAD
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
=======
	af, err := os.Open(filenames[0])
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}
	defer af.Close()
	a, err := ioutil.ReadAll(af)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}

	bf, err := os.Open(filenames[1])
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}
	defer bf.Close()
	b, err := ioutil.ReadAll(bf)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}

	fmt.Fprint(c.outStream, diff(string(a), string(b), unified))
>>>>>>> upstream/master

	return 0
}
