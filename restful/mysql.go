package main
import (
    "database/sql"
    "fmt"
    "log"
   // _ "github.com/go-sql-driver/mysql"
    _ "github.com/Go-SQL-Driver/MySQL")
func main() {
    //open*******************************
    db, err := sql.Open("mysql", "jchjjj:jiangchuan@/gotest") //jchjjj:jiangchuan@
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    //create table *******************************
    /*	sql := `
    create table bar (id int not null primary key, name text) ;
    `
    _, err = db.Exec(sql)
    if err != nil {
        log.Printf("%q: %s\n", err, sql)
        return
    } else{
        //		log.Printf("create table bar success!")
    }
    */
    //insert***************************
    sql := `insert into test values('a','a'),('b','b'),('c','c');`
    _,err = db.Exec(sql)
    checkErr(err)
    //select ***************************
    sql = `select * from test;`
    data,err := db.Query(sql)
    for data.Next() {
        var name string
        var passwd string
        err = data.Scan(&name,&passwd)
        fmt.Println(name,passwd)
    }
    // update*****************************
    stmt, err := db.Prepare("update test set passwd=? where name=?")
    checkErr(err)
    res, err := stmt.Exec("aab", "a")
    checkErr(err)
    affect, err := res.RowsAffected()
    fmt.Println(affect)
    checkErr(err)
    stmt.Close()
    //delete ****************************
    stmt,err = db.Prepare("delete from test where name=?")
    checkErr(err)
    res,err = stmt.Exec("a")
    affect,err = res.RowsAffected()
    fmt.Println(affect,"rows deleted!")
}
func checkErr(err error){
    if err != nil{
        panic(err)
    }
}
