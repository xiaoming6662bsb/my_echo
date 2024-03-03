package utils

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func GetByteNanoTime(t int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t))
	return buf
}
func DeCodeByteNanoTime(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
func Nano2Time(na int64) time.Time {
	seconds := na / int64(time.Second)     // 计算秒
	nanoseconds := na % int64(time.Second) // 计算纳秒的剩余部分
	return time.Unix(seconds, nanoseconds)
}
func GetLocal() *time.Location {
	location, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		fmt.Println("GetLocal Err:", err)
		return nil
	}
	return location
}

type NtpPacket struct {
	Settings          byte
	Stratum           byte
	Poll              byte
	Precision         byte
	RootDelay         uint32
	RootDispersion    uint32
	ReferenceID       uint32
	RefTimestamp      uint64
	OrigTimestamp     uint64
	RecvTimestamp     uint64
	TransmitTimestamp uint64
}

// utils.GetSync("1.debian.pool.ntp.org:123")
func GetSync(serv string) {
	addr, err := net.ResolveUDPAddr("udp", serv)
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer conn.Close()

	// 初始化NTP请求包
	request := &NtpPacket{Settings: 0x1B} // 设置模式为3（客户端）和版本号为3

	// 发送请求
	err = binary.Write(conn, binary.BigEndian, request)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}

	// 接收服务器响应
	response := &NtpPacket{}
	err = binary.Read(conn, binary.BigEndian, response)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	// 计算当前时间
	secs := float64(response.TransmitTimestamp >> 32)
	frac := float64(response.TransmitTimestamp&0xFFFFFFFF) / 0x100000000
	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	currentTime := ntpEpoch.Add(time.Duration(secs * float64(time.Second))).Add(time.Duration(frac * float64(time.Second)))

	fmt.Println("Server Time:", currentTime.UnixNano())
	fmt.Println("Local  Time:", time.Now().UnixNano())
}
