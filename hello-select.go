package main


import (
    "fmt"
    "time"
    )
    
    
func fibonacci(c, quit chan int){
	x,y := 1,1
	for {
		select {
		case c <- x:
			x,y = y,x+y
		case <- quit:
			fmt.Println("quit")
		case <- time.After(5* time.Second):
			fmt.Println("timeout")
		    break
		default:
			fmt.Println("default")
		}
	}
}
   

    
func main(){
	c := make(chan int)
	quit := make(chan int)
	go func(){
		for i:= 0;i < 10; i++ {
			fmt.Println(<-c)	
		}
		quit <- 3
	}()
	fibonacci(c,quit)
	
}
