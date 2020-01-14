package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//返回两个返回值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	//获取方法名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sumArgs(values ...int) int {
	sum := 0
	fmt.Println(values)
	for i := range values {
		sum += values[i]
	}
	return sum
}

//交换a,b位置
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(2, 3, "_"))
	q, r := div(5, 3)
	fmt.Println(q, r)

	fmt.Println(apply(
		func(i int, i2 int) int {
			return int(math.Pow(float64(i), float64(i2)))
		}, 2, 3))

	fmt.Println(apply(pow, 2, 3))
	fmt.Println("1+2+3+4+5=", sumArgs(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
