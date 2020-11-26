package main

import "testing"


/*using interface
it's not necessary that the class declare implemention

the type can implement a method in the interface
when the method is used referring the interface,
the compiler create autmatically a link
between the type and the interface
doubt: how about a method with same name and more than one interface?"
**/

/**
table driven tests
*/


func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
         {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {shape: Circle{Radius: 10}, want: 314.1592653589793},
        {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	/*/clearer than 
	{Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
        {Triangle{12, 6}, 36.0},
	*/
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}







