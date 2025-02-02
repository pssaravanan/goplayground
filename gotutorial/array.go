package gotutorial

import "fmt"

func ArrayMain(){
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	var s1 = arr[1:4]
	var s2 = arr[3:4]
	fmt.Println("Array:")
	fmt.Println("\t Array intialization", arr)
	fmt.Println("\t Slice1", s1)
	fmt.Println("\t Slice2", s2)
	s2[0] = 100
	fmt.Println("\t After Modification:")
	fmt.Println("\t Slice1", s1)
	fmt.Println("\t Slice2", s2)
	s3 := []int{1, 2, 3}
	fmt.Println("\t Slice3 initialization", s3)
	fmt.Println("\t Slice3 s3[1:]", s3[1:])
	fmt.Println("\t length of s3", len(s3))
	fmt.Println("\t cap of s3", cap(s3))
	fmt.Println("\t length of s3[1:]", len(s3[1:]))
	fmt.Println("\t cap of s3[1:]", cap(s3[1:]))
	fmt.Println("\t cap of s2", cap(s2))
	fmt.Println("\t s2", s2)


	var s4 = [][]string{
		[]string {"a", "b"},
		[]string {"d", "e", "f"},
	}

	fmt.Println("\t slices of slices:", s4)

	s2 = append(s2, 200)
	fmt.Println("\t slice after append", s2)

	fmt.Println("\t Ranges")
	for i, v  := range(s2) {
		fmt.Printf("\t i: %v, v: %v\n", i, v)
	}
}