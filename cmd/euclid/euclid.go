package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/sauterp/euclid"
)

var usage string = `Usage: euclid a b
non-negative integer a < 18446744073709551615
non-negative integer b < 18446744073709551615
`

func printErrExit(err string) {
	fmt.Fprintf(os.Stderr, "%s\n\n%s\n", err, usage)
	os.Exit(1)
}

func checkInputArg(arg string) uint64 {
	i, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		printErrExit(fmt.Sprintf("Error: %s\nunable to parse argument: %s\n", err.Error(), arg))
		os.Exit(1)
	}
	return i
}

func main() {

	if len(os.Args) < 3 {
		printErrExit("Error: too few arguments")
	}

	a := checkInputArg(os.Args[1])
	b := checkInputArg(os.Args[2])

	d, err := euclid.GCD(a, b)
	if err != nil {
		printErrExit(err.Error())
	}
	fmt.Println(d)

}
