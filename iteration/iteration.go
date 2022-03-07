// To do stuff repeatedly in Go, you'll need for. In Go there are no while, do, until keywords, you can only use for. Which is a good thing!
package iteration

func Repeat(char string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}
	return repeated
}
