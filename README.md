Telnet implements basic support for telnet protocol.

WARNING: This is a work in progress, so the API can change at any point.
If you want stability, vendor it into your own repository.

Example how to use:

``` go
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
```