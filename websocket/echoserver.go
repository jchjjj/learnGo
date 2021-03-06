/* EchoServer
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	// "io"
	"code.google.com/p/go.net/websocket"
)

func Echo(ws *websocket.Conn) {
	fmt.Println(ws)
	fmt.Println(ws.RemoteAddr())
	fmt.Println("Echoing")
	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+48)
		fmt.Println("Sending to client: " + msg)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("Can't send")
			break
		}
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":12345", nil)
	//err := http.ListenAndServeTLS(":12345","jiang","chuan",nil)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
