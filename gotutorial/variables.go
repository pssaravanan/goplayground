package gotutorial

var x, y int

const GO = "GO"

func LocalVariable()(z int){
	z = 10
	return
}

func PackageVariable()(int){
	y = 10
	return y
}

func PackageVariableReadOnly()(int){
	return y
}

func MultipleVariables()(x, y, z int){
	x, y, z = 10, 20, 30
	return
}

func VariableTypes()(int, string, float64, string, bool){
	var (
		a int
		b string
		c float64 = 30.0
		d string = "Hello"
	)

	return a, b, c, d, true

}


func VariableTypeConversion(a int)(float64){
	return float64(a)
}