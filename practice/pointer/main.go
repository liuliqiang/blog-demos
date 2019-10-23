package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	slice(a)

	b := [][]int{}
	fmt.Printf("before: %p\n", &b)
	twoSlice(b)
	fmt.Printf("after: %p\n", &b)

	c := [][]int{}
	twoSlicePointer(&c)

	d := make([]int, 10)
	singleSlice(d)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
}


func slice(a []int) {
	a[1] = 3
}
func singleSlice(a []int) {
	 a= append(a, 1)
}

func twoSlice(b [][]int) {
	fmt.Printf("in before: %p\n", &b)
	b = append(b, []int{1,2, 3})
	fmt.Printf("in after: %p\n", &b)
}
func twoSlicePointer(b *[][]int) {
	*b = append(*b, []int{1,2, 3})
}