package main
import "fmt"

func main(){
    fmt.Printf("hello 世界！\n")
    const (
        x = iota
        y = iota
        z = iota
    )
    fmt.Printf("%v,%v,%v",x,y,z)

    var arr[10]int
    arr[1]=21
    fmt.Printf("%d\n",arr[1])
    bo := true
    fmt.Printf("bool:%v\n",bo)
    var m map[string] int
    m = make(map[string] int)
    m["a"] = 1
    m["haha"] = 3333333
    fmt.Printf("map:%v %v\n",m["a"],m["haha"])

    if x:= add(1,2);x>10 {
        fmt.Printf("larger than 10\n")
    }else{
        fmt.Printf("smaller than 10\n")
    }
    
    sum := 0
    for i := 0 ;i<1000000000;i++{
        sum += i
    }
    fmt.Printf("sum : %x\n",sum)
    println(sum)
}
func add(a int ,b int) int{
    return a+b
}
