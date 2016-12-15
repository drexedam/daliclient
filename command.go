package daliclient

// MakeBroadcastCmd creates a command to be broadcasted to all devices
// arg is only needed for
// CmdSetScene
// CmdStoreScene
// CmdRemoveScene
// CmdAddGroup
// CmdRemoveGroup
// CmdScene
func MakeBroadcastCmd(arg byte, command Command) []byte {
	return append(
		[]byte{
			daliVersion,
			0x00,
		},
		makeCommand(arg, 0xff, command)...,
	)
}

// MakeLampCmd creates a command to be sent to the lamp given by dest
// arg is only needed for
// CmdSetScene
// CmdStoreScene
// CmdRemoveScene
// CmdAddGroup
// CmdRemoveGroup
// CmdScene
func MakeLampCmd(arg byte, dest byte, command Command) []byte {
	return makeCommonCmd(arg, dest, 0x1, command)
}

// MakeGroupCmd creates a command to be sent to the group given by dest
// arg is only needed for
// CmdSetScene
// CmdStoreScene
// CmdRemoveScene
// CmdAddGroup
// CmdRemoveGroup
// CmdScene
func MakeGroupCmd(arg byte, dest byte, command Command) []byte {
	return makeCommonCmd(arg, dest, 0xf, command)
}

// MakeSpecialCommand creates a special command
// cmd is only needed for
// DestDtr
// DestInit
// DestSearchh
// DestSearchm
// DestSearchl
// DestSet
// DestCheck
func MakeSpecialCommand(cmd byte, dest Destination) []byte {
	var code0 byte
	var code1 byte
	switch dest {
	case DestTerm:
		code0 = 0xa1
		code1 = 0x00
	case DestDtr:
		code0 = 0xa3
		code1 = cmd & 0xff
	case DestInit:
		code0 = 0xa5
		code1 = cmd & 0xff
	case DestRandom:
		code0 = 0xa7
		code1 = 0x00
	case DestCompare:
		code0 = 0xa9
		code1 = 0x00
	case DestWithdraw:
		code0 = 0xab
		code1 = 0x00
	case DestSearchh:
		code0 = 0xb1
		code1 = cmd & 0xff
	case DestSearchm:
		code0 = 0xb3
		code1 = cmd & 0xff
	case DestSearchl:
		code0 = 0xb5
		code1 = cmd & 0xff
	case DestSet:
		code0 = 0xb7
		code1 = cmd & 0xff
	case DestCheck:
		code0 = 0xb9
		code1 = cmd & 0xff
	case DestAddress:
		code0 = 0xbb
		code1 = 0x00
	case DestPhys:
		code0 = 0xbd
		code1 = 0x00
	case DestDefault:
		fallthrough
	default:
		code0 = 0x00
		code1 = 0x00
	}

	return []byte{code0, code1}
}

// makeCommonCmd creates the complete command including address and version
func makeCommonCmd(arg, dest, and byte, command Command) []byte {
	address := makeAddress(dest, and)
	return append(
		[]byte{
			daliVersion,
			0x00,
		},
		makeCommand(arg, address, command)...,
	)
}

// makeAddress creates an address
func makeAddress(dest, and byte) byte {
	return 0x1 | ((dest & and) << 1)
}

// makeCommand creates the requested command
func makeCommand(arg byte, address byte, command Command) []byte {
	var code byte
	switch command {
	case CmdOff:
		code = 0x00
	case CmdDimUp:
		code = 0x01
	case CmdDimDown:
		code = 0x02
	case CmdInc:
		code = 0x03
	case CmdDec:
		code = 0x04
	case CmdSetMax:
		code = 0x05
	case CmdSetMin:
		code = 0x06
	case CmdDown:
		code = 0x07
	case CmdUp:
		code = 0x08
	case CmdSetScene:
		code = 0x10 | (arg | 0xf)
	case CmdReset:
		code = 0x20
	case CmdStoreDtr:
		code = 0x21
	case CmdStoreMax:
		code = 0x2a
	case CmdStoreMin:
		code = 0x2b
	case CmdStoreFail:
		code = 0x2c
	case CmdStorePower:
		code = 0x2d
	case CmdStoreTime:
		code = 0x2e
	case CmdStoreRate:
		code = 0x2f
	case CmdStoreScene:
		code = 0x40 | (arg | 0xf)
	case CmdRemoveScene:
		code = 0x50 | (arg | 0xf)
	case CmdAddGroup:
		code = 0x60 | (arg | 0xf)
	case CmdRemoveGroup:
		code = 0x70 | (arg | 0xf)
	case CmdStoreAddress:
		code = 0x80
	case CmdStatus:
		code = 0x90
	case CmdCheckWork:
		code = 0x91
	case CmdCheckLamp:
		code = 0x92
	case CmdCheckOperat:
		code = 0x93
	case CmdCheckLevel:
		code = 0x94
	case CmdCheckReset:
		code = 0x95
	case CmdCheckAddress:
		code = 0x96
	case CmdVersion:
		code = 0x97
	case CmdDtr:
		code = 0x98
	case CmdType:
		code = 0x99
	case CmdPhysical:
		code = 0x9a
	case CmdCheckFail:
		code = 0x9b
	case CmdLevel:
		code = 0xa0
	case CmdMax:
		code = 0xa1
	case CmdMin:
		code = 0xa2
	case CmdPower:
		code = 0xa3
	case CmdFail:
		code = 0xa4
	case CmdTimeRate:
		code = 0xa5
	case CmdScene:
		code = 0xb0 | (arg | 0xf)
	case CmdGroup0:
		code = 0xc0
	case CmdGroup8:
		code = 0xc1
	case CmdRandomh:
		code = 0xc2
	case CmdRandomm:
		code = 0xc3
	case CmdRandoml:
		code = 0xc4
	case CmdDefault:
		fallthrough
	default:
		code = 0x0

	}

	return []byte{address, code}
}
