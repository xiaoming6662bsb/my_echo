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

	timestamp := time.Now().In(utils.GetLocal())
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
	stime := utils.Nano2Time(utils.DeCodeByteNanoTime(buffer))
	etime := time.Now().In(utils.GetLocal())

	fmt.Printf("\n\r\n\033[34mTCP\u001B[0m start\n%v\nup   \033[32m[%v]\033[0m\n%v\ndown \033[32m[%v]\033[0m\n%v\nall  \033[32m[%v]\033[0m\n\033[34mTCP\033[0m end\n\r\n",
		timestamp.UnixNano(),
		stime.Sub(timestamp).String(),
		stime.UnixNano(),
		etime.Sub(stime).String(),
		etime.UnixNano(),
		etime.Sub(timestamp).String(),
	)
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

	timestamp := time.Now().In(utils.GetLocal())
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
		stime := utils.Nano2Time(utils.DeCodeByteNanoTime(byt))
		etime := time.Now().In(utils.GetLocal())

		fmt.Printf("\n\r\n\033[33mUDP\u001B[0m start\n%v\nup   \033[32m[%v]\033[0m\n%v\ndown \033[32m[%v]\033[0m\n%v\nall  \033[32m[%v]\033[0m\n\033[33mUDP\033[0m end\n\r\n",
			timestamp.UnixNano(),
			stime.Sub(timestamp).String(),
			stime.UnixNano(),
			etime.Sub(stime).String(),
			etime.UnixNano(),
			etime.Sub(timestamp).String(),
		)
		return
	}
}
