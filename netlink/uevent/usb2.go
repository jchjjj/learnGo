package main

import (
	"log"
	"os"
	"strings"
	"syscall"
)

const NETLINK_KOBJECT_UEVENT = 15
const UEVENT_BUFFER_SIZE = 2048

func main() {
	fd, err := syscall.Socket(
		syscall.AF_NETLINK, syscall.SOCK_RAW,
		NETLINK_KOBJECT_UEVENT,
	)
	if err != nil {
		log.Println(err)
		return
	}

	nl := syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pid:    uint32(os.Getpid()),
		Groups: 1,
	}
	err = syscall.Bind(fd, &nl)
	if err != nil {
		log.Println(err)
		return
	}

	b := make([]byte, UEVENT_BUFFER_SIZE*2)
	for {
		n, err2 := syscall.Read(fd, b)
		log.Println(n, err2)
        log.Println(b)
        act,dev,subsys := parseUBuffer(b)
        log.Println(act,dev,subsys)
	}
}
func parseUBuffer(arr []byte) (act, dev, subsys string) {
	j := 0
	for i := 0; i < len(arr)+1; i++ {
		if i == len(arr) || arr[i] == 0 {
			str := string(arr[j:i])
			a := strings.Split(str, "=")
			if len(a) == 2 {
				switch a[0] {
				case "DEVNAME":
					dev = a[1]
				case "ACTION":
					act = a[1]
				case "SUBSYSTEM":
					subsys = a[1]
				}
			}
			j = i + 1
		}
	}

	return

}
