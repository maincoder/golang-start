package main


import (
    "fmt"
    "runtime"
    )
    
func say(s string,count int) {
	for i := 0 ;i < count;i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}
    
    
func main(){
    go say("Hello",4)
    say("world",2)
}

