/* func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}


What we need to do is to be able to inject (which is just a fancy word for pass in) the dependency of printing.
Our function doesn't need to care _where** or **how** the printing happens, so we should accept an **interface_** rather than a concrete type.
*/

package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
