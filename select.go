package main
import (
    "runtime"
    "fmt"
)

func fibonacci(c,quit chan int){
    x,y := 1,1
    for {
        select{
        case c <- x:
            x,y = y,x+y
        case <-quit :
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    var  a  =[]int{1,2,3,4,5,6}
    c := make(chan int)
    quit := make(chan int)
    go fibonacci(c,quit)
    for i:= range a {
        fmt.Println(i,<-c)
    }
    fmt.Println("cpu nums:",runtime.NumCPU())
    fmt.Println("num go routines:",runtime.NumGoroutine())
    fmt.Println("Go max procs:",runtime.GOMAXPROCS(0))
    quit <- 0
}
