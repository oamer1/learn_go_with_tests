package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// spy sleeper
type sleeper interface {
	Sleep()
}

// Spies are a kind of mock which can record how a dependency is used
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// io.Writer is general
func Countdown(out io.Writer) {

	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)

	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}
func main() {
	Countdown(os.Stdout)
}
