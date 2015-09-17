package main
import (
"fmt"
"time"

)
func Sleep(x int){
	<-time.After(time.Second*time.Duration(x))
}


func main() {
	t1:= time.Now()
	Sleep(10)
	t2:=time.Now()
fmt.Println(t1,t2)
}