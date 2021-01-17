package main

import (
	"fmt"
)

type geoShape interface {
	getArea() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width float64
}

func (c *Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (r *Rectangle) getArea() float64 {
	return r.length * r.width
}	
	
func printArea(shape geoShape) float64 {
	return shape.getArea()
}	
 
func main(){
	c := Circle{4.5}
	r := Rectangle{1.2, 3.6}
	fmt.Println(c.getArea())
	fmt.Println(r.getArea())
	
	shapes := []geoShape{&c, &r} // slice of structs
	for _, shape := range shapes {
		fmt.Println("in slices")
		fmt.Println(shape.getArea())
		fmt.Println("in func")
		fmt.Println(printArea(shape))
	}
		 
}
// good practice to pass by ref instead of val
