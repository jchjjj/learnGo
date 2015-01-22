/* PersonClientJSON
 */
package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	var msg string
	service := os.Args[1]
	origin := "http://localhost"
	conn, err := websocket.Dial(service, "", origin)
	checkError(err)
	person := Person{Name: "Jan",
		Emails: []string{"ja@newmarch.name", "jan.newmarch@gmail.com"},
	}
	for {
		err = websocket.JSON.Send(conn, person)
		if err != nil {
			fmt.Println("Couldn't send msg " + err.Error())
		}
		err = websocket.Message.Receive(conn, &msg)
		checkError(err)
		fmt.Println("recv from server: ", msg)
	}

	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			fmt.Println("end of file\n")
			os.Exit(1)
		} else {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(1)
		}

	}
}
