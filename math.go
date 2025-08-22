package math

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Div(a, b int) int {
	return a / b
}
func Mul(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Add(5, 3)) //sample
}
