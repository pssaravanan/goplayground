package gotutorial

import "fmt"

func MapMain() {
	sampleMap := make(map[string]int)
	fmt.Println("Map main")
	fmt.Println("\t nil map", sampleMap)

	var currencyShort = map[string]string {
		"US Dollar": "USD",
		"UK Pound": "GBP",
		"India Rupees": "INR",
	}
	var nilMap map[int]int = make(map[int]int)
	fmt.Println("\t rupees short form", currencyShort)
	fmt.Println("\t nilMap", nilMap)

	nilMap[1] = 10
	fmt.Println("\t nilMap", nilMap)

	elem, ok := nilMap[1]
	fmt.Println("\t elem", elem)
	fmt.Println("\t ok", ok)

	delete(nilMap, 1)
	elem, ok = nilMap[1]
	fmt.Println("\t elem", elem)
	fmt.Println("\t ok", ok)
}