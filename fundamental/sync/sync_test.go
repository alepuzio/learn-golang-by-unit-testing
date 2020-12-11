package main

import (
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}
/**
* I pass Counter to avoid error message:
*** call of assertCounter copies lock value (by doc: A Mutex must not be copied after first use.)

*/
func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()//Helper marks the calling function as a test helper function
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}









