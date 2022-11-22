package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Reader struct {
}

func (r *Reader) read() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Fprint(conn, "GET HTTP 1.0\r\n\r\n")

	readerToStdout(conn, 100)
}
func fileReader() {
	f, err := os.Open("a")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	readerToStdout(f, 10)
}

func readerToStdout(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
