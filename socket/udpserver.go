package main 

import (
	"fmt"
	"net"
	"os"
)

func checkerr(err error) {
	if err != nil {
		fmt.Printf("error :%v",err.Error())
		os.Exit(1)
	}
}
func main() {
	var buf [1500]byte
	udpAddr,err := net.ResolveUDPAddr("udp4",":9899")
	conn ,err:=net.ListenUDP("udp",udpAddr)
	checkerr(err)
	for {
		n,_,err := conn.ReadFromUDP(buf[0:])
		checkerr(err)
		if n>0 {
//			msg := string(buf[0:n])
			//fmt.Println(msg)
		}
	}
}
