package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var port = flag.Int("port", 8000, "Port number used for clock server.")
	var tz = flag.String("tz", "US/Eastern", "RFC defined timezone")

	flag.Parse()

	err := os.Setenv("TZ", *tz)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Starting clock for %s on port %d\n", *tz, *port)

	host := "localhost:" + strconv.Itoa(*port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// handle incoming connections concurrently
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
