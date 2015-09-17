package main 
import "testing"

func TestFibSeries(t *testing.T) {

	var num int
	num = fib(5)
	if num != 5 {
		t.Error("Expected 5, got", num)
	}

}

