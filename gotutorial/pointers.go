package gotutorial

import "fmt"
func PointersValueChange(x *int) {
	*x = 20
}

func PointersMain(){
	fmt.Println("Pointers")
	a := 10
	b := &a
	fmt.Println("\t Variable", a)
	fmt.Println("\t Pointer Value", b)
	PointersValueChange(b)
	fmt.Println("\t Changed Pointer value", a)

}