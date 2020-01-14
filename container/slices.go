package main

import "fmt"

func updateSilce(arr []int) {
	arr[0] = 100
}

func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr1[2:6]
	s2 := arr1[:]
	fmt.Println("s1 =", s1)
	fmt.Println("arr1[2:] =", arr1[2:])
	fmt.Println("arr1[:6] =", arr1[:6])
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSilce(s1)")
	updateSilce(s1)
	fmt.Println("s1 =", s1)
	fmt.Println(arr1)

	fmt.Println("Extending slice")
	arr1[0], arr1[2] = 0, 2
	fmt.Println("arr1 =", arr1)
	s1 = arr1[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5 =", s3, s4, s5)
	fmt.Println(arr1)
}
