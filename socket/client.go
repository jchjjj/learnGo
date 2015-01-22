package main 

import (
	"fmt"
	"net"
	"os"
)

var str string
var msg = make([]byte ,1024)
func main() {
	var(
		host	=	"127.0.0.1"
		port	=	"9897"
		remote	=	host+":"+port
		//data	=	make([]byte,1024)
	)

	con,err := net.Dial("tcp",remote)
	defer con.Close()
	if err != nil {
		fmt.Println("sever not found.")
		os.Exit(-1)
	}
	fmt.Println("connection OK.")

	for {
		fmt.Printf("Enter a sentence:")	
		fmt.Scanf("%s",&str)
		if str == "quit" {
			fmt.Println("communication terminated.")
			os.Exit(-1)
		}

		_,err := con.Write([]byte(str))
		if err != nil {
			fmt.Println("error when send to server.")
			os.Exit(-1)
		}

		length,err := con.Read(msg)
		if err != nil {
			fmt.Println("error when read from server.")
			os.Exit(-1)
		}

		str = string(msg[0:length])
		fmt.Println(str)
	}

}