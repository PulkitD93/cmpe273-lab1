package main

import (
   "fmt"
   "math"
)

/* define an interface */
type Shape interface {
   perimeter() float64
}

type Circle struct { 
   rad float64
}


type Rectangle struct {
   width, height float64
}

func(circle Circle) perimeter() float64 {// Method for perimeter of a circle
   return 2*math.Pi * circle.rad
}

func(rect Rectangle) perimeter() float64 {//Method for perimeter of Rectangle
   return 2*(rect.width + rect.height)
}

func main() {
   circle := Circle{10}
   rectangle := Rectangle {10, 5}
   
   fmt.Printf("Circle perimeter: %f\n",circle.perimeter())
   fmt.Printf("Rectangle perimeter: %f\n",rectangle.perimeter())
}