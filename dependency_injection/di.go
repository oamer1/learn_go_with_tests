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
	"log"
	"net/http"
)

// As discussed earlier fmt.Fprintf allows you to pass in an io.Writer
// which we know both os.Stdout and bytes.Buffer implement.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter also implements io.Writer so this is why we could re-use our Greet
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

// When you write an HTTP handler, you are given an http.ResponseWriter and the http.Request that was used to make the request. When you implement your server you write your response using the writer.
