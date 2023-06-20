package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "0.0.0.0:54321"

const proto = "tcp4"

var quotes = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {

	rand.Seed(time.Now().UnixNano())

	listener, err := net.Listen(proto, addr)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go sendRandomQuote(conn)
	}

}

func sendRandomQuote(conn net.Conn) {

	ticker := time.NewTicker(3 * time.Second)

	defer conn.Close()

	defer ticker.Stop()

	

	for {
		select {
		case <-ticker.C:
			conn.Write([]byte(takeRandomQuote() + "\n"))
		}
	}

}

func takeRandomQuote() string {

	index := rand.Intn(len(quotes))
	return quotes[index]
}
