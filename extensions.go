package telnet

import (
	"fmt"
	"sync"
)

// InitNegotiator creates a new Negotiator for a particular code
type InitNegotiator func(Command) Negotiator

// Extensions implements simple multi-negotiator,
// it will automatically create negotiator and dispatch to the appropriate one
type Extensions struct {
	Initiators   map[Option]InitNegotiator // constructors for negotiators
	Negotiations map[Option]Negotiator     // stores any ongoing negotations
}

func NewExtensions(initiators map[Option]InitNegotiator) Negotiator {
	return &Extensions{
		Initiators:   initiators,
		Negotiations: make(map[Option]Negotiator),
	}
}

func (n *Extensions) Handle(c Command) Command {
	cmd, ok := c.(OptionCommand)
	if !ok {
		return nil
	}
	opt := cmd.OptionCode()

	ongoing, ok := n.Negotiations[opt]
	if ok {
		return ongoing.Handle(cmd)
	}

	init, ok := n.Initiators[opt]
	if ok {
		neg := init(cmd)
		n.Negotiations[opt] = neg
		return neg.Handle(cmd)
	}

	return Wont{opt}
}

var (
	registryMu sync.Mutex
	registry   = make(map[Option]InitNegotiator)
)

// Register registers a new negotiator constructor
func Register(option Option, init InitNegotiator) {
	registryMu.Lock()
	defer registryMu.Unlock()

	if _, registered := registry[option]; registered {
		panic(fmt.Sprintf("option code %d already registered", option))
	}
	registry[option] = init
}

// RegisteredNegotiators returns a map containing all registered negotiators
func RegisteredNegotiators() map[Option]InitNegotiator {
	registryMu.Lock()
	defer registryMu.Unlock()

	r := make(map[Option]InitNegotiator, len(registry))
	for code, init := range registry {
		r[code] = init
	}
	return r
}
