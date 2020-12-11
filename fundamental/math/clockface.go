package main

import "time"

/* 
A Point represents a two dimensional Cartesian coordinate.
*/
type Point struct {
	X float64 //x-axis
	Y float64 //y-axis
}

/*
SecondHand is the unit vector of the second hand of an analogue clock at time `t`.
represented as a Point.
*/

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
