package main 

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkerr(err error) {
	if err != nil {
		fmt.Printf("error :%v",err.Error())
		os.Exit(1)
	}
}
func main() {
	var buf = make([]byte,468)
    if len(os.Args)<2 {
        fmt.Printf("address and port should be given\n")
        os.Exit(1)
    }
	remoteAddr := os.Args[1]
	port := os.Args[2]
	udpAddr,err := net.ResolveUDPAddr("udp4",remoteAddr+":"+port)
	conn ,err:=net.DialUDP("udp",nil,udpAddr)
	checkerr(err)
	for {
		_,err := conn.Write([]byte(buf))
		checkerr(err)
		time.Sleep(time.Nanosecond*500)
	}

}
