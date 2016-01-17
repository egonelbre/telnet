package telnet

import (
	"fmt"
	"io"
	"strings"
)

const maxLine = 4096

// Negotiator handles any negotiations to the telnet protocol
type Negotiator interface {
	Handle(c Command) (response Command)
}

// Conn handles reading/writing telnet protocol and can handle extensions
type Conn struct {
	conn       io.ReadWriter
	Lines      chan string
	send       chan Command
	negotiator Negotiator
}

func NewConn(input io.ReadWriter) *Conn {
	return NewConnWithNegotiator(input, NewExtensions(RegisteredNegotiators()))
}

func NewConnWithNegotiator(input io.ReadWriter, n Negotiator) *Conn {
	r := &Conn{
		conn:       input,
		Lines:      make(chan string),
		send:       make(chan Command),
		negotiator: n,
	}
	go r.run()
	return r
}

func (r *Conn) formatLine(line string) string {
	return strings.Replace(line, "\n", "\r\n", -1)
}

func (r *Conn) prepare(c Command) Command {
	switch cmd := c.(type) {
	case string:
		return r.formatLine(cmd)
	case Transaction:
		for i, c := range cmd {
			cmd[i] = r.prepare(c)
		}
		return cmd
	default:
		return cmd
	}
}

func (r *Conn) Send(c Command) {
	r.send <- r.prepare(c)
}

func (r *Conn) Print(s string) {
	r.send <- r.prepare(s)
}

func (r *Conn) Printf(format string, a ...interface{}) {
	r.send <- r.prepare(fmt.Sprintf(format, a...))
}

func (r *Conn) run() {
	defer r.Terminate()

	cmds := make(chan Command)
	go Unserialize(r.conn, cmds)
	go Serialize(r.send, r.conn)

	for c := range cmds {
		switch cmd := c.(type) {
		default:
			// log.Printf("Unexpected type %T\n", cmd)
		case string:
			// log.Printf("User action '%s'\n", cmd)
			r.Lines <- cmd
		case OptionCommand:
			// info := cmd.OptionCode().Info()
			// log.Printf("Command [%v] %v\n", info.Name, cmd)
			if r.negotiator != nil {
				r.negotiator.Handle(cmd)
			} else {
				r.Send(Wont{cmd.OptionCode()})
			}
		}
	}
}

func (r *Conn) Terminate() {
	close(r.Lines)
}
