package main


import (
    "fmt"
//    "time"
    )
    
    
func fibonacci(c, quit chan int){
	x,y := 1,1
	for {
		select {
		case c <- x:
			x,y = y,x+y
		case <- quit:
			fmt.Println("quit")
//		case <- time.After(5* time.Second):
//			fmt.Println("timeout")
		    return
//		default:
//			fmt.Println("default")
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
	m := map[string]string{"key1":"val1"}
	fmt.Println(m)
	m["key2"]="val2"
	fmt.Println(m)
	delete(m,"key1")
	fmt.Println(m)
	
	m_int := map[string]int {"key1":3}
	fmt.Println(m_int)
	m_int["key2"]=4
	fmt.Println(m_int)
	delete(m_int,"key1")
	
	
	fmt.Println(m_int)
	m_float := map[int]float32 {3:2.5}
	m_float[7] = 2.0
	fmt.Println(m_float)
	
	
	
}
