package telnet

import "fmt"

// This is a listing of different Telnet RFC-s

type Option byte

func (opt Option) Info() OptionInfo { return GetOptionInfo(opt) }

type OptionInfo struct {
	Code     Option
	Mnemonic string
	Name     string
}

func GetOptionInfo(opt Option) OptionInfo {
	if opt, ok := bycode[opt]; ok {
		return opt
	}
	return OptionInfo{opt, fmt.Sprintf("U%d"), fmt.Sprintf("Unknown (%d)")}
}

var bycode = map[Option]OptionInfo{
	OptionBINARY:         {OptionBINARY, "BINARY", "binary"},
	OptionECHO:           {OptionECHO, "ECHO", "echo"},
	OptionRCP:            {OptionRCP, "RCP", "prepare for reconnect"},
	OptionSGA:            {OptionSGA, "SGA", "supress go ahead"},
	OptionNAMS:           {OptionNAMS, "NAMS", "approximate message size"},
	OptionSTATUS:         {OptionSTATUS, "STATUS", "give status"},
	OptionTIMING_MARK:    {OptionTIMING_MARK, "TIMING_MARK", "timing mark"},
	OptionRCTE:           {OptionRCTE, "RCTE", "remote controlled transmission and echo"},
	OptionNAOL:           {OptionNAOL, "NAOL", "negotiate about ouput line width"},
	OptionNAOP:           {OptionNAOP, "NAOP", "negotiate about output page size"},
	OptionNAOCRD:         {OptionNAOCRD, "NAOCRD", "negotiate about CR disposition"},
	OptionNAOHTS:         {OptionNAOHTS, "NAOHTS", "negotiate about horizontal tabstops"},
	OptionNAOHTD:         {OptionNAOHTD, "NAOHTD", "negotiate about horizontal tab disposition"},
	OptionNAOFFD:         {OptionNAOFFD, "NAOFFD", "negotiate about formfeed disposition"},
	OptionNAOVTS:         {OptionNAOVTS, "NAOVTS", "negotiate about vertical tab stops"},
	OptionNAOVTD:         {OptionNAOVTD, "NAOVTD", "negotiate about vertical tab disposition"},
	OptionNAOLFD:         {OptionNAOLFD, "NAOLFD", "negotiate about output LF disposition"},
	OptionXASCII:         {OptionXASCII, "XASCII", "extended ascii character set"},
	OptionLOGOUT:         {OptionLOGOUT, "LOGOUT", "force logout"},
	OptionBM:             {OptionBM, "BM", "byte macro"},
	OptionDET:            {OptionDET, "DET", "data entry terminal"},
	OptionSUPDUP:         {OptionSUPDUP, "SUPDUP", "supdup protocol"},
	OptionSUPDUPOUTPUT:   {OptionSUPDUPOUTPUT, "SUPDUPOUTPUT", "supdup output"},
	OptionSNDLOC:         {OptionSNDLOC, "SNDLOC", "send location"},
	OptionTTYPE:          {OptionTTYPE, "TTYPE", "terminal type"},
	OptionEOR:            {OptionEOR, "OPT_EOR", "end or record"},
	OptionTUID:           {OptionTUID, "TUID", "TACACS user identification"},
	OptionOUTMRK:         {OptionOUTMRK, "OUTMRK", "output marking"},
	OptionTTYLOC:         {OptionTTYLOC, "TTYLOC", "terminal location number"},
	OptionVT3270REGIME:   {OptionVT3270REGIME, "VT3270REGIME", "3270 regime"},
	OptionX3PAD:          {OptionX3PAD, "X3PAD", "X.3 PAD"},
	OptionNAWS:           {OptionNAWS, "NAWS", "window size"},
	OptionTERMINAL_SPEED: {OptionTERMINAL_SPEED, "TERMINAL_SPEED", "terminal speed"},
	OptionLFLOW:          {OptionLFLOW, "LFLOW", "remote flow control"},
	OptionLINE_MODE:      {OptionLINE_MODE, "LINE_MODE", "Linemode option"},
	OptionXDISPLOC:       {OptionXDISPLOC, "XDISPLOC", "X Display Location"},
	OptionOLD_ENVIRON:    {OptionOLD_ENVIRON, "OLD_ENVIRON", "Old - Environment variables"},
	OptionAUTHENTICATION: {OptionAUTHENTICATION, "AUTHENTICATION", "Authenticate"},
	OptionENCRYPT:        {OptionENCRYPT, "ENCRYPT", "Encryption option"},
	OptionNEW_ENVIRON:    {OptionNEW_ENVIRON, "NEW_ENVIRON", "Environment variables"},
	OptionCHARSET:        {OptionCHARSET, "CHARSET", "Charset"},
	OptionMSDP:           {OptionMSDP, "MSDP", "Mud Server Data Protocol"},
	OptionMSSP:           {OptionMSSP, "MSSP", "Mud Server Status Protocol"},
	OptionMCCPv1:         {OptionMCCPv1, "MCCPv1", "Mud Client Compression Protocol v1"},
	OptionMCCP:           {OptionMCCP, "MCCP", "Mud Client Compression Protocol v2"},
	OptionMSP:            {OptionMSP, "MSP", "MUD Sound Protocol"},
	OptionMXP:            {OptionMXP, "MXP", "Mud eXtension Protocol"},
	OptionZMP:            {OptionZMP, "ZMP", "Zenith MUD Protocol"},
	OptionGCMP:           {OptionGCMP, "GCMP", "Generic Mud Communication Protocol"},
	OptionEXOPL:          {OptionEXOPL, "EXOPL", "telnet extended options"},
}

const (
	OptionBINARY         Option = 0  // RFC 856
	OptionECHO           Option = 1  // RFC 857
	OptionRCP            Option = 2  // RFC 426
	OptionSGA            Option = 3  // RFC 858
	OptionNAMS           Option = 4  //
	OptionSTATUS         Option = 5  // RFC 859
	OptionTIMING_MARK    Option = 6  // RFC 860
	OptionRCTE           Option = 7  //
	OptionNAOL           Option = 8  //
	OptionNAOP           Option = 9  //
	OptionNAOCRD         Option = 10 //
	OptionNAOHTS         Option = 11 //
	OptionNAOHTD         Option = 12 //
	OptionNAOFFD         Option = 13 //
	OptionNAOVTS         Option = 14 //
	OptionNAOVTD         Option = 15 //
	OptionNAOLFD         Option = 16 //
	OptionXASCII         Option = 17 //
	OptionLOGOUT         Option = 18 //
	OptionBM             Option = 19 //
	OptionDET            Option = 20 //
	OptionSUPDUP         Option = 21 //
	OptionSUPDUPOUTPUT   Option = 22 //
	OptionSNDLOC         Option = 23 //
	OptionTTYPE          Option = 24 // RFC 930, 1091, http://tintin.sourceforge.net/mtts/
	OptionEOR            Option = 25 // RFC 885
	OptionTUID           Option = 26 //
	OptionOUTMRK         Option = 27 //
	OptionTTYLOC         Option = 28 //
	OptionVT3270REGIME   Option = 29 //
	OptionX3PAD          Option = 30 //
	OptionNAWS           Option = 31
	OptionTERMINAL_SPEED Option = 32  // RFC 1079
	OptionLFLOW          Option = 33  //
	OptionLINE_MODE      Option = 34  // RFC 1184
	OptionXDISPLOC       Option = 35  //
	OptionOLD_ENVIRON    Option = 36  //
	OptionAUTHENTICATION Option = 37  //
	OptionENCRYPT        Option = 38  //
	OptionNEW_ENVIRON    Option = 39  // RFC 1572
	OptionCHARSET        Option = 42  // RFC 2066
	OptionMSDP           Option = 69  // http://tintin.sourceforge.net/msdp/
	OptionMSSP           Option = 70  // http://tintin.sourceforge.net/mssp/
	OptionMCCPv1         Option = 85  // http://www.zuggsoft.com/zmud/mcp.htm
	OptionMCCP           Option = 86  // http://www.zuggsoft.com/zmud/mcp.htm
	OptionMSP            Option = 90  // http://www.zuggsoft.com/zmud/msp.htm
	OptionMXP            Option = 91  // http://www.zuggsoft.com/zmud/mxp.htm
	OptionZMP            Option = 93  // http://discworld.starturtle.net/external/protocols/zmp.html
	OptionGCMP           Option = 201 // http://www.ironrealms.com/gmcp-doc
	OptionEXOPL          Option = 255 // RFC 861
)
