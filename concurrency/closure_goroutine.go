// https://go.dev/doc/faq#closures_and_goroutines
// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
// https://stackoverflow.com/questions/21822527/why-golang-dont-iterate-correctly-in-my-for-loop-with-range
package concurrency

import (
	"fmt"
)

func con() {
	vals := []int{1, 2, 3}
	done := make(chan bool)

	for _, v := range vals {
		go func() {
			fmt.Println(v)
			done <- true

		}()

	}

	// wait for coroutines to exit
	for _ = range vals {
		<-done
	}

}

/* output
3
3
3

each iteration of the loop uses the same instance of the variable v,
so each closure shares that single variable. When the closure runs,
it prints the value of v at the time fmt.Println is executed, but v may
 have been modified since the goroutine was launched. To help detect this and
 other problems before they happen, run go vet

or

for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
In this example, the value of v is passed as an argument to the anonymous function. That value is then accessible inside the function as the variable u.

Even easier is just to create a new variable, using a declaration style that may seem odd but works fine in Go:

    for _, v := range values {
        v := v // create a new 'v'.
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }


because the goroutines will probably not begin executing until after the loop.


*/
