package main

import (
	"fmt"
	"io"
	"os"
)

// io.Writer is general
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
func main() {
	Countdown(os.Stdout)
}
