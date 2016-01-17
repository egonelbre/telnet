package ttype

import (
	"github.com/egonelbre/telnet"
)

// Implements the TTYPE protocol as defined by RFC 1091
// http://www.ietf.org/rfc/rfc1091.txt

const Code = telnet.OptionTTYPE

type Negotiation struct {
	Type string
}

func New() *Negotiation {
	return &Negotiation{""}
}

func init() {
	telnet.Register(Code,
		func(c telnet.Command) telnet.Negotiator { return New() })
}

func (n *Negotiation) Request() telnet.Command {
	return telnet.Do{Code}
}

func (n *Negotiation) Handle(c telnet.Command) telnet.Command {
	switch cmd := c.(type) {
	case telnet.Will:
		// client showed willingness
		return telnet.SubNegotiation{Code, []byte{telnet.SEND}}
	case telnet.SubNegotiation:
		n.Type = string(cmd.Data[1:])
		return nil
	}
	return nil
}
