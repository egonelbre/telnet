package telnet

type (
	// Command represents any element that can be serialized/unserialized
	// by telnet protocol
	Command interface{}
	// OptionCommand
	OptionCommand interface {
		Command
		OptionCode() Option
	}

	Break            struct{}
	InterruptProcess struct{}
	AbortOutput      struct{}
	AreYouThere      struct{}
	GoAhead          struct{}

	Will struct{ Option Option }
	Wont struct{ Option Option }
	Do   struct{ Option Option }
	Dont struct{ Option Option }

	SubNegotiation struct {
		Option Option
		Data   []byte
	}

	// for sending only
	Transaction []Command
)

func (c Will) OptionCode() Option           { return c.Option }
func (c Wont) OptionCode() Option           { return c.Option }
func (c Do) OptionCode() Option             { return c.Option }
func (c Dont) OptionCode() Option           { return c.Option }
func (c SubNegotiation) OptionCode() Option { return c.Option }
