package server

import (
	"fmt"
	"meo/internal/utils"
	"net"
	"time"
)

func handleTCPConn(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	buffer := make([]byte, 8)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	t := time.Now().In(utils.GetLocal())

	_, err = conn.Write(utils.GetByteNanoTime(t.UnixNano()))
	if err != nil {
		fmt.Println("Error Write:", err)

		return
	}
	stime := utils.Nano2Time(utils.DeCodeByteNanoTime(buffer))

	fmt.Printf("\n\r\nTCP[%v]\n%v\nin[%v]\n%v\nTCP\n\r\n",
		conn.RemoteAddr().String(),
		stime.UnixNano(),
		t.Sub(stime).String(),
		t.UnixNano(),
	)

}

func StartTCP(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer func() {
		_ = listener.Close()
	}()
	fmt.Println("TCP on", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleTCPConn(conn)
	}
}

func StartUDP(addr string) {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	fmt.Println("UDP on", addr)

	buffer := make([]byte, 8)
	for {
		_, laddr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}
		t := time.Now().In(utils.GetLocal())
		_, err = conn.WriteTo(utils.GetByteNanoTime(t.UnixNano()), laddr)
		if err != nil {
			fmt.Println("Error WriteTo from connection:", err)
		}
		stime := utils.Nano2Time(utils.DeCodeByteNanoTime(buffer))

		fmt.Printf("\n\r\nUDP[%v]\n%v\nin[%v]\n%v\nUDP\n\r\n",
			laddr.String(),
			stime.UnixNano(),
			t.Sub(stime).String(),
			t.UnixNano(),
		)
	}
}
