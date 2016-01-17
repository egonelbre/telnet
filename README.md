Telnet implements basic support for telnet protocol.

Example how to use:

``` go
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/egonelbre/telnet"
)

func main() {
	fmt.Printf("Server started on :6000\n")
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

const greeting = `

    #################
    #               #
    #  HELLO WORLD  #
    #               #
    #################

`

func handleConnection(c net.Conn) {
	r := telnet.NewConn(c)

	r.Print(greeting)
	r.Print("What's your nick? ")
	nick := <-r.Lines

	r.Printf("\033[1;30;41mWelcome %s!\033[0m\n", nick)

	for line := range r.Lines {
		fmt.Printf("...doing %s\n", line)
	}
	fmt.Printf("finished!!!\n")
}
```