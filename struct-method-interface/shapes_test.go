package main

import "testing"


/*using interface
it's not necessary that the class declare implemention

the type can implement a method in the interface
when the method is used referring the interface,
the compiler create autmatically a link
between the type and the interface
doubt: how about a method with same name and more than one interface?"
func TestArea(t *testing.T) {

    checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    }

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        checkArea(t, rectangle, 72.0)
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        checkArea(t, circle, 314.1592653589793)
    })

}
**/

/**
table driven tests
*/
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        print(got)
	if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}













