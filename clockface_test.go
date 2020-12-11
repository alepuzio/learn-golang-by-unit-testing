package main

import (
	"testing"
	"time"
	"github.com/quii/learn-go-with-tests/math/v1/clockface"
)
/*
before go test, run 

go get github.com/quii/learn-go-with-tests/math/v1/clockface

to download the clockface object 
*/

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
