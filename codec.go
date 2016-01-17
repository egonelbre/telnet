package telnet

import (
	"bufio"
	"io"
	"log"
)

const maxLineLength = 1024

type PreprocessLine func(string) string

func write(c Command, to *bufio.Writer) {
	switch cmd := c.(type) {
	case string:
		to.WriteString(cmd)
	case []byte:
		to.Write(cmd)

	case Transaction:
		for _, c := range cmd {
			write(c, to)
		}

	case SubNegotiation:
		to.Write([]byte{IAC, SB, byte(cmd.Option)})
		to.Write(cmd.Data)
		to.Write([]byte{IAC, SE})

	case Will:
		to.Write([]byte{IAC, WILL, byte(cmd.Option)})
	case Wont:
		to.Write([]byte{IAC, WONT, byte(cmd.Option)})
	case Do:
		to.Write([]byte{IAC, DO, byte(cmd.Option)})
	case Dont:
		to.Write([]byte{IAC, DONT, byte(cmd.Option)})

	case Break:
		to.Write([]byte{IAC, BRK})
	case InterruptProcess:
		to.Write([]byte{IAC, IP})
	case AbortOutput:
		to.Write([]byte{IAC, AO})
	case AreYouThere:
		to.Write([]byte{IAC, AYT})
	case GoAhead:
		to.Write([]byte{IAC, GA})
	}
}

// Serializes Commands into io.Writer
func Serialize(in chan Command, to io.Writer) {
	stream := bufio.NewWriter(to)
	for c := range in {
		write(c, stream)
		stream.Flush()
	}
}

// Unserializes telnet bytestream into commands
func Unserialize(input io.Reader, out chan Command) {
	stream := bufio.NewReader(input)

	next := func() byte {
		b, err := stream.ReadByte()
		if err != nil {
			panic(err)
		}
		return b
	}

	defer func() {
		err := recover()
		if (err == io.EOF) || (err == nil) {
			return
		}

		log.Printf("Error occurred: %v\n", err)
	}()

	buffer := make([]byte, 0, maxLineLength)
	for {
		b := next()
		switch b {
		case IAC:
			command := next()
			switch command {
			case SB:
				sub := SubNegotiation{Option(next()), make([]byte, 0)}
				for {
					b := next()
					if b == IAC {
						suboption := next()
						if suboption == SB {
							sub.Data = append(sub.Data, IAC)
						} else if suboption == SE {
							break
						} else {
							panic("Protocol error!")
						}
					} else {
						sub.Data = append(sub.Data, b)
					}
				}
				out <- sub
			case WILL:
				out <- Will{Option(next())}
			case WONT:
				out <- Wont{Option(next())}
			case DO:
				out <- Do{Option(next())}
			case DONT:
				out <- Dont{Option(next())}

			case BRK:
				out <- Break{}
			case IP:
				out <- InterruptProcess{}
			case AO:
				out <- AbortOutput{}
			case AYT:
				out <- AreYouThere{}

			case EC:
				if len(buffer) > 0 {
					buffer = buffer[:len(buffer)-1]
				}
			case EL:
				buffer = buffer[:0]

			case DM:
				panic("Don't know how to handle data mark!")

			case NOP, GA: // do nothing
			default:
				buffer = append(buffer, command)
			}
		case CarriageReturn, LineFeed:
			if len(buffer) > 0 {
				out <- string(buffer)
				buffer = buffer[:0]
			}
		default:
			buffer = append(buffer, b)
			if len(buffer) >= maxLineLength {
				panic("Line too long!")
			}
		}
	}
}
