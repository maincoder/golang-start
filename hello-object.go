package main


import (
    "fmt"
    "math"
    )
    
    
type Rectangle struct {
	width,height float64
}    

type Circle struct{
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle)area() float64 {
	return c.radius * c.radius * math.Pi
}

    
func main(){
    fmt.Println("Hello World!")
    r1 := Rectangle{12,2}
    r2 := Rectangle{9,4}
    c1 := Circle{10}
    c2 := Circle{9}
    fmt.Println("Area of r1 is: ",r1.area())
    fmt.Println("Area of r1 is: ",r2.area())
    fmt.Println("Area of c1 is: ",c1.area())
    fmt.Println("Area of c1 is: ",c2.area())
}

