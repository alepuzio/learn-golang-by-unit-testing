package main

import (
	"fmt"
	"io"
	"time"
)

// Countdown prints a countdown from 3 to out.


const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
    for i := countdownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
    }
    
    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}
