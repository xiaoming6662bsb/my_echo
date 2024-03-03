package utils

import (
	"encoding/binary"
	"fmt"
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
