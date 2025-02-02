package main

import (
	"fmt"
	"time"

	// "github.com/pssaravanan/goplayground/lc"
	"github.com/pssaravanan/goplayground/gotutorial"
)

func main() {
	// result := lc.SortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})
	// fmt.Println(result)
	// lc.ReverseParentheses("(ed(et(oc))el)")
	// lc.SortJumbled([]int{8,9,4,0,2,1,3,5,7,6}, []int{991,338,38})
	// result := lc.SortArray([]int{-2,3,-5})
	fmt.Println("Start:", time.Now())
	gotutorial.HelloWorld()
	fmt.Println("Functions:")
	sum := gotutorial.Sum(10, 20)
	fmt.Println("\t Sum:", sum)
	a, b := gotutorial.Swap(10, 20)
	fmt.Println("\t Swap:", a, b)
	fmt.Println("Variables:")
	fmt.Println("\t Local Variable", gotutorial.LocalVariable())
	fmt.Println("\t Package Variable", gotutorial.PackageVariable())
	fmt.Println("\t Package Variable Readonly", gotutorial.PackageVariableReadOnly())
	a, b, c := gotutorial.MultipleVariables()
	fmt.Println("\t Package Multiple Variables", a, b, c)
	a1, b1, c1, d1, e1 := gotutorial.VariableTypes()
	fmt.Println("\t Package Variables Types", a1, b1, c1, d1, e1)
	fmt.Printf("\t Package Variables Types %T, %T, %T, %T, %T \n", a1, b1, c1, d1, e1)
	fmt.Printf("\t Package Variables Format values %v, %v, %v, %v, %v \n", a1, b1, c1, d1, e1)
	fmt.Println("\t Package Variables Type conversion", gotutorial.VariableTypeConversion(10))
	fmt.Println("\t Package Variables Constant", gotutorial.GO)
	fmt.Println("FlowControl:")
	gotutorial.BasicLoop()
	gotutorial.LoopWithOnlyCondition()
	gotutorial.LoopsWithoutConditions()
	fmt.Println("\t IfEven", gotutorial.IfEven(11))
	fmt.Println("\t Switch", gotutorial.SwitchCase(1))
	fmt.Println("\t Switch Variables", gotutorial.SwitchCaseWithoutVariable(10))
	gotutorial.DeferSum(10, 20)

	fmt.Println("Pointers:")
	gotutorial.PointersMain()

	gotutorial.StructMain()

	gotutorial.ArrayMain()

	gotutorial.MapMain()

	gotutorial.MethodMain()
}
