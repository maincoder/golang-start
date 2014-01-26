package main


import (
    "fmt"
    )
    
    
func sum(a []int, c chan int){
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}
   

    
func main(){
	c := make(chan int,3)
	c <- 2
	c <- 1
	fmt.Println(<-c)
	fmt.Println(<-c)
}
