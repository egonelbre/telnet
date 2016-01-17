package negotiator

import (
	"github.com/egonelbre/telnet/opt"
	"github.com/egonelbre/telnet/opt/naws"
	"github.com/egonelbre/telnet/opt/ttype"
)

var DefaultInitiator = map[byte]Initiator{
	opt.NAWS:  naws.InitByClient,
	opt.TTYPE: ttype.InitByClient,
}
