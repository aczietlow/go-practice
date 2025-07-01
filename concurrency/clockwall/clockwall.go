package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	clocks := []int{8000, 8001}

	for _, port := range clocks {
		go newConn("localhost:" + strconv.Itoa(port))
	}
	for {
		// literally wait forever
	}

}
func newConn(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
