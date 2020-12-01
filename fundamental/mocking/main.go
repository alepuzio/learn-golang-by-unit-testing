package main

import (
	"fmt"
	"io"
)

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer) {
    for i := 3; i > 0; i-- {
        fmt.Fprintln(out, i)
    }
    fmt.Fprint(out, "Go!")
}
