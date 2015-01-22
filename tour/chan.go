package main

import "fmt"

func sum(start, end int64, c chan int64) (res int64) {
	for i := start; i < end; i++ {
		res += i
	}
	c <- res
	return res
}
func main() {
	var total int64 = 1 << 34
	/*
		c := make(chan int64)
		go sum(0,total/4,c)
		go sum(total/4+1,total/2,c)
		go sum(total/2+1,total*3/4,c)
		go sum(total*3/4+1,total,c)
		s := make([]int64,4)
		var sum int64 = 0
		for i:=0;i<4;i++{
			s[i] = <-c
			fmt.Println(s[i])
			sum += s[i]
		}
	*/
	var sum, i int64 = 0, 0
	for i = 0; i < total; i++ {
		sum += i
	}
	fmt.Println(sum)
}
