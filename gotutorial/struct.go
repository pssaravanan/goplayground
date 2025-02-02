package gotutorial

import "fmt"
import "math"

type Point struct {
	X int
	Y int
}

func DistanceBetweenPoints(x, y Point) float64 {
	return math.Sqrt(float64((x.X - y.X) * (x.X - y.X) + (x.Y - y.Y) * (x.Y - y.Y)))
}

func StructMain(){
	fmt.Println("Struct:")
	x := Point{X: 10, Y: 20}
	y := Point{X: 15, Y: 25}
	fmt.Println("\t Distance:", DistanceBetweenPoints(x, y))
}