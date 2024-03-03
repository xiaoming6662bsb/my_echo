package main

import (
	"flag"
	"fmt"
	"meo/internal/age"
	"meo/internal/server"
	"sync"
)

func main() {
	var isAge bool
	var isServer bool
	var TcpPort string
	var UcpPort string
	var RemoteHost string
	flag.BoolVar(&isAge, "age", false, "Run as age client")
	flag.BoolVar(&isServer, "serv", false, "Run as server")
	flag.StringVar(&TcpPort, "tp", "8080", "server tcp port")
	flag.StringVar(&UcpPort, "up", "8081", "server udp port")
	flag.StringVar(&RemoteHost, "th", "", "server ip")
	flag.Parse()

	if isAge && isServer {
		fmt.Println("Cannot run as both age and server")
		return
	}
	if (RemoteHost == "" || len(RemoteHost) == 0) && isAge {
		fmt.Println("Cannot RemoteHost empty and server")
		return
	}
	if isAge {
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go age.SendTCP(RemoteHost+":"+TcpPort, wg)
		go age.SendUDP(RemoteHost+":"+UcpPort, wg)
		wg.Wait()
	} else if isServer {
		go server.StartTCP("0.0.0.0:" + TcpPort)
		go server.StartUDP("0.0.0.0:" + UcpPort)

		st := make(chan struct{}, 1)
		<-st
	} else {
		fmt.Println("Please specify --age or --serv")
	}
}
