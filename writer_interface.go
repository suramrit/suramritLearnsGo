package main 


import("fmt"
		"os"
		"io"
)

// Fromt the docs :: 
// type Writer interface {
//         Write(p []byte) (n int, err error)
// }

// Writer is the interface that wraps the basic Write method.

// Write writes len(p) bytes from p to the underlying data stream. 
// -- Suramrit: without the need for moving it into a programs DS -- think of file ops/ network ops etc. 
// It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered 
// that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p). 
// Write must not modify the slice data, even temporarily.

// Many imp packages like os, fmt, file, etc implement write... so can be used in methods like Fprinln, encoder, decoder as well 

// For more details look at the golang documentation - https://golang.org/pkg/io/#Writer
func main() {
	fmt.Println("Hey there!")
	fmt.Fprintln(os.Stdout, "Hola!!")
	io.WriteString(os.Stdout, "Hola again!! -- so many ways because of interfaces! \n")
}













