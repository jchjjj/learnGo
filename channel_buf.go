package main
import "fmt"
func main() {
    c := make(chan int, 1)// 修改2 为1 就报错,修改2 为3 可以正常运行
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)

}
