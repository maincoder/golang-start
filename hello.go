package main


import (
    "fmt"
    )
    
    
type Rectangle struct {
	width,height float64
}    

func area(r Rectangle) float64 {
	return r.width * r.height
}

    
func main(){
    fmt.Println("Hello World!")
    var c complex64 = 5+5i
    var d complex64 = 5+6i
    fmt.Println(c*d)
}

