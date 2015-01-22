package main

import (
    "fmt"
    "github.com/astaxie/beego/httplib"
)
func main(){
    /*defer func(){
        recover()
        fmt.Println("recover from panic")
    }()
    */
    //post**********************
    req := httplib.Post("http://localhost:8080/users")
    for i :=1; i<10000; i++ {
        //req.Param("id",fmt.Sprintf("%d",i))
        //req.Param("name",fmt.Sprintf("name%d",i))
      //  fmt.Println(i)
       // req.Param("id","1")
       //req.Param("{\"name\":\"jac\"}")
        //req.Param("name","hehe")
        req.Body(fmt.Sprintf("{\"name\":\"jac%d\"}",i))//for json ,must use Body method,key=value is not right
       result,err := req.String()
        if err != nil {
            fmt.Println("error")
            fmt.Println(err)
        }else{
            fmt.Println(result)
        }
    }
    /*
    str, err := httplib.Get("http://localhost:8080/users").String()
    if err != nil {
         //t.Fatal(err)
         fmt.Println(err)
         //panic("error in httplib")
    }
    fmt.Println(str)
    */
}    
