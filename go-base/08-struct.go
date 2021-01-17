package main

import (
	"fmt"
)

type student struct {
	name string
	age int
} 

type Point struct {
	x float64
	y float64
}

type Circle struct {
	r float64
	center *Point // same as only *Point
}

func changeX(pt *Point) { // pass by value pt Point
	pt.x = 5.6
	fmt.Println(pt)
}



func main(){
	s := student{"abc xyz", 10}
	fmt.Println(s)
	fmt.Println(s.age)
	
	var p Point = Point{1.5, 2.3}
	fmt.Printf("%f, %f \n", p.x, p.y)
	
	pTest := Point{y : 3.5}
	fmt.Println(pTest)  // {0 3.5}
		    
	pTest1 := &Point{y : 3.5}
	fmt.Println(pTest1)  // &{0 3.5}
	
	pTest2 := &Point{y : 3.5} // pass by ref
	fmt.Println(pTest2)  // &{0 3.5}
	changeX(pTest2)
	fmt.Println(pTest2)  // &{5.6 3.5}
		    
	pTest3 := Point{y : 3.5} // pass by value
	fmt.Println(pTest3)  // {0 3.5}
	//changeX(pTest3)
	fmt.Println(pTest3)  // {0 3.5}	    -> don't do it
			    
	c := Circle{4.56, pTest2}		    
	fmt.Println(c)	// {4.56 0xc00002c090}
	fmt.Println(c.center.x) // 5.6 // same as c.x but better to define a var in struct 
	
		    
}

