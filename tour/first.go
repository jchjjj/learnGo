package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	v.X = 4
	fmt.Println(v.X, v.Y)
	a := [2]Vertex{}
	a[0] = Vertex{1, 2}
	for index, _ := range a {
		a[index] = Vertex{index, index}
	}
	for index, val := range a {
		fmt.Printf("a[%d]={%d,%d}\n", index, val.X, val.Y)
		fmt.Printf("a[%d]=%v\n", index, val)
	}
	fmt.Println(a[0].X, a[0].Y)
	//b :=make( []Vertex,8)
	b := []Vertex{{1, 2}, {2, 3}}
	fmt.Println("b is", b, len(b), cap(b))
	b = a[:1]
	fmt.Println(b[0].X, b[0].Y)
	fmt.Println("b is", b, len(b), cap(b))

	c := new(Vertex)
	c = &a[0]
	fmt.Println(c)
	c.X = 1000
	fmt.Println(a[0].X, a[0].Y)
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",
			i, p[i])
	}
	//function is value too
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))
	//cloure
	plus, minux := adder(), adder()
	for i := 1; i < 10; i++ {
		fmt.Println(plus(i), minux(-i))
	}
}

func adder() func(int) int {
	sum := 0 // sum is bound to the function returned
	return func(x int) int {
		sum += x
		return sum
	}
}
