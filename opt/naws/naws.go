package naws

import (
	"github.com/egonelbre/telnet"
)

// Implements the NAWS protocol as defined by RFC 1073
// http://www.ietf.org/rfc/rfc1073.txt

const Code byte = 31

type Negotiation struct {
	ServerAsked bool
	Width       int
	Height      int
}

func New() *Negotiation {
	return &Negotiation{true, -1, -1}
}

func InitByClient(c telnet.Command) telnet.Negotiator {
	return &Negotiation{false, -1, -1}
}

func (n *Negotiation) Request() telnet.Command {
	return telnet.Do{Code}
}

func (n *Negotiation) Handle(c telnet.Command) telnet.Command {
	switch cmd := c.(type) {
	case telnet.Will:
		// if the server didn't ask, we need to respond
		if !n.ServerAsked {
			return telnet.Do{Code}
		}
	case telnet.SubNegotiation:
		n.Width = int(cmd.Data[0])<<4 + int(cmd.Data[1])
		n.Height = int(cmd.Data[2])<<4 + int(cmd.Data[3])
		return nil
	}
	return nil
}
