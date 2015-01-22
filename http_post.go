package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "log"
)
var count int = 0
func sayhelloName(w http.ResponseWriter, r * http.Request){
    count++
    fmt.Println(count)
    r.ParseForm();
    fmt.Println(r.Form)
    fmt.Println("path",r.URL.Path)
    fmt.Println("scheme",r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k,v := range r.Form {
        fmt.Println("key:",k)
        fmt.Println("value:",strings.Join(v,""))
    }
    fmt.Fprintf(w,"hello go http")
    //add a new page
    /*
    http.HandleFunc("/haha",func(w http.ResponseWriter,r * http.Request){
        fmt.Fprintf(w,"haha,you got me!")
    })
    */
}
func login(w http.ResponseWriter,r *http.Request){
    r.ParseForm();
    fmt.Println("method:",r.Method)
    if r.Method == "GET" {
        t,_:= template.ParseFiles("login.gtpl")
        t.Execute(w,nil)
    } else {
        fmt.Println("username:",r.Form["username"])
        fmt.Println("password:",r.Form["password"])
    }
}
func main(){
    http.HandleFunc("/",sayhelloName)
    http.HandleFunc("/login",login)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("listenandserve: ",err)
    }
}
