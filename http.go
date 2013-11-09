package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)
var count int = 0;
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
}

func main(){
    http.HandleFunc("/",sayhelloName)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("listenandserve: ",err)
    }
}
