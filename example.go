// +build ignore

package main

import (
	"fmt"

	"github.com/egonelbre/telnet"
)

func main() {
	fmt.Printf("Server started on :8000\n")
	telnet.ListenAndServe(":8000", handle)
}

func handle(conn *telnet.Conn) {
	conn.Print("\n\n#  HELLO WORLD  #\n\n")
	conn.Print("What's your nick? ")
	nick := <-conn.Lines

	conn.Printf("Welcome %s!\n", nick)

	for line := range conn.Lines {
		fmt.Printf("[%s] %s\n", nick, line)
	}
}
