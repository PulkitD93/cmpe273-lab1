package main
import "fmt"
func fib(n int) int {
	if n==0{
		return 0
	}else if n==1{
		return 1
		}else{
	return fib(n-1) + fib(n-2)

}}
 func main() {
	fmt.Print("Enter a Number")
	var x int
	fmt.Scanf("%d",&x)
	fmt.Print(fib(x))
}