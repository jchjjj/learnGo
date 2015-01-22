package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var(
		host	=	"127.0.0.1"
		port	=	"9897"
		remote	=	host+":"+port
		data	=	make([]byte,1024)
	)

	fmt.Println("initiating server ,press Ctrl+C to stop")

	lis,err	:=	net.Listen("tcp",remote)
	defer lis.Close()
	if err != nil {
		fmt.Println("error when Listen:",remote)
		os.Exit(-1)
	}

	for  {
		var res string
		conn,err := lis.Accept()
		if err != nil {
			fmt.Println("errro in Accept:",err.Error())
			os.Exit(0)
		}

		go func(con net.Conn) {
			fmt.Println("new connection:",con.RemoteAddr())
			for{
				length,err := con.Read(data)
				if err != nil {
					fmt.Printf("Client %v quit.\n",con.RemoteAddr)
					con.Close()
					return
				}
			
				res = string(data[0:length])
				fmt.Printf("%s said:%s\n",con.RemoteAddr(),res)
				res = "you said:"+res
				con.Write([]byte (res))
			}
		}(conn)
	}
}