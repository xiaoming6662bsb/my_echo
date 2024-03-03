package age

import (
	"fmt"
	"meo/internal/utils"
	"net"
	"sync"
	"time"
)

func SendTCP(serverAddr string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	timestamp := time.Now()
	_, err = conn.Write(utils.GetByteNanoTime(timestamp.UnixNano()))
	if err != nil {
		fmt.Println("Error Write to server:", err)
		return
	}

	buffer := make([]byte, 8)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	etime := time.Now()
	stime := utils.Nano2Time(utils.DeCodeByteNanoTime(buffer))

	fmt.Printf("\n\r\nTCP\n%v\nup[%v]\n%v\ndown[%v]\n%v\nTCP\n\r\n", timestamp.UnixNano(), stime.Sub(timestamp).String(), stime.UnixNano(), etime.Sub(stime).String(), etime.UnixNano())
}

func SendUDP(serverAddr string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	timestamp := time.Now()
	go func() {
		_, _ = conn.Write(utils.GetByteNanoTime(timestamp.UnixNano()))
	}()
	for {
		byt := make([]byte, 8)
		_, err = conn.Read(byt)
		if err != nil {
			fmt.Println("UDP Error Read", err)
			return
		}
		etime := time.Now()
		stime := utils.Nano2Time(utils.DeCodeByteNanoTime(byt))
		fmt.Printf("\n\r\nUDP\n%v\nup[%v]\n%v\ndown[%v]\n%v\nUDP\n\r\n", timestamp.UnixNano(), stime.Sub(timestamp).String(), stime.UnixNano(), etime.Sub(stime).String(), etime.UnixNano())
		return
	}
}
