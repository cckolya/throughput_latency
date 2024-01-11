package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

var (
	// кол-во циклов нужно, чтобы нагрузить CPU
	cpuTimes = map[string]int{
		"1ms":   3000,
		"5ms":   12500,
		"10ms":  25000,
		"20ms":  50000,
		"300ms": 1000000,
	}
)

// cpu input key from cpuTimes
func cpu(time string) error {
	countFor, exist := cpuTimes[time]
	if !exist {
		return errors.New("такого времени нет, введите 1ms, 5ms, 10ms, 20ms, 300ms")
	}
	h := md5.New()
	io.WriteString(h, "Hello, world!")
	for i := 0; i < countFor; i++ {
		s := fmt.Sprintf("%x", h.Sum(nil))
		io.WriteString(h, s)
		// io.WriteString(os.Stdout, s + "\n")
	}
	return nil
}

//func main() {
//	var rusageStart, rusageStop syscall.Rusage
//	var start time.Time
//	start = time.Now()
//	syscall.Getrusage(syscall.RUSAGE_SELF, &rusageStart)
//
//	h := md5.New()
//	io.WriteString(h, "Hello, world!")
//	for i := 0; i < 3000; i++ {
//		s := fmt.Sprintf("%x", h.Sum(nil))
//		io.WriteString(h, s)
//		// io.WriteString(os.Stdout, s + "\n")
//	}
//
//	time.Sleep(2 * time.Second)
//	syscall.Getrusage(syscall.RUSAGE_SELF, &rusageStop)
//	delta := time.Since(start)
//	fmt.Printf("total time = %+v\n", delta)
//
//	delta = time.Duration(rusageStop.Utime.Sec)*time.Second +
//		time.Duration(rusageStop.Utime.Usec)*time.Microsecond -
//		time.Duration(rusageStart.Utime.Sec)*time.Second -
//		time.Duration(rusageStart.Utime.Usec)*time.Microsecond
//
//	fmt.Printf("user time = %+v\n", delta)
//
//	delta = time.Duration(rusageStop.Stime.Sec)*time.Second +
//		time.Duration(rusageStop.Stime.Usec)*time.Microsecond -
//		time.Duration(rusageStart.Stime.Sec)*time.Second -
//		time.Duration(rusageStart.Stime.Usec)*time.Microsecond
//
//	fmt.Printf("system time = %+v\n", delta)
//}
