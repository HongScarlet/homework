package temperature

import "fmt"

// C to F
func C2F(x float32) {
	var temp float32
	temp=x*1.8+32
	fmt.Printf("%f C = %f F\n", x, temp)
}

// F to C
func F2C(x float32) {
	var temp float32
	temp=(x-32)/1.8
	fmt.Printf("%f F = %f C\n", x, temp)
}
