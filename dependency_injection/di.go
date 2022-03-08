/* func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}


What we need to do is to be able to inject (which is just a fancy word for pass in) the dependency of printing.
Our function doesn't need to care _where** or **how** the printing happens, so we should accept an **interface_** rather than a concrete type.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

// As discussed earlier fmt.Fprintf allows you to pass in an io.Writer
// which we know both os.Stdout and bytes.Buffer implement.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
