package gotutorial

import "fmt"

func BasicLoop(){
	fmt.Printf("\t Basic Loop ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i)
	}
	fmt.Printf("\n")
}

func LoopWithOnlyCondition(){
	fmt.Printf("\t Loop with only condition ")
	i := 0
	for i < 10 {
		fmt.Printf("%v", i)
		i++
	}
	fmt.Printf("\n")
}

func LoopsWithoutConditions(){
	fmt.Printf("\t Loop without condition ")
	i := 0
	for {
		fmt.Printf("%v", i)
		i++
		if i == 10 {
			break
		}
	}
	fmt.Printf("\n")
}

func IfEven(x int) bool {
	if result := x % 2; result == 0 {
		return true
	}
	return false
}

func SwitchCase(x int) string {
	var result string
	switch x {
	case 1:
		result = "1"
	case 2:
		result = "2"
	default:
		result = ">2"
	}
	return result
}

func SwitchCaseWithoutVariable(x int) string {
	var result string
	switch {
	case x == 1:
		result = "1"
	case x == 2:
		result = "2"
	default:
		result = ">2"
	}
	return result
}

func DeferSum(x, y int) int{
	defer fmt.Println("\t Deffered Sum:", Sum(x, y))
	fmt.Println("\t Addition")
	return 10
}