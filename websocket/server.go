package main

import (
	"archive/zip"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)
		for i := 0; i < 10; i++ {
			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
			time.Sleep(time.Microsecond * 500000)
		}

	}
}
func client(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<meta charset="UTF-8">
<html>
<head></head>
<body>
<script type="text/javascript">
var sock = null;
var wsuri = "ws://127.0.0.1:12345";
window.onload = function() {
console.log("onload");
try{
	sock = new WebSocket(wsuri);	
}catch(e){
	alert(e.Message);
}

sock.onopen = function() {
console.log("connected to " + wsuri);
}
sock.onclose = function(e) {
	console.log("connection closed (" + e.code + ")");
}
sock.onmessage = function(e) {
console.log("message received: " + e.data);
}
};
function send() {
var msg = document.getElementById('message').value;
sock.send(msg);
};
</script>
<h1>WebSocket Echo Test</h1>
<form>
<p>
Message: <input id="message" type="text" value="Hello, world!">
</p>
</form>
<button onclick="send();">Send Message</button>
</body>
</html>

	`
	io.WriteString(w, html)
}
func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.HandleFunc("/chat", client)
	fmt.Println("listen on port 12345.")
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
