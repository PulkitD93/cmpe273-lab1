package main 
import "testing"

func TestShapeInterface(t *testing.T) {

	var periCircle, periRect float64
	circle := Circle{ 10}
	rect := Rectangle{ 10, 5}

	periCircle = circle.perimeter()
	periRect = rect.perimeter()

	if periCircle != 62.83185307179586 {
		t.Error("Expected 62.83185307179586, got", periCircle)
	}

	if periRect != 30 {
		t.Error("Expected 30 ,got", periRect)
	}

}
