package daliclient

// Destination for special commands
type Destination int

// Command is an enum for the available commands
type Command int

// Only needed for special commands
const (
	DestDefault Destination = iota
	DestTerm
	DestDtr
	DestInit
	DestRandom
	DestCompare
	DestWithdraw
	DestSearchh
	DestSearchm
	DestSearchl
	DestSet
	DestCheck
	DestAddress
	DestPhys
)

// Command to send
const (
	CmdDefault Command = iota
	CmdOff
	CmdDimUp
	CmdDimDown
	CmdInc
	CmdDec
	CmdSetMax
	CmdSetMin
	CmdDown
	CmdUp
	CmdSetScene
	CmdReset
	CmdStoreDtr
	CmdStoreMax
	CmdStoreMin
	CmdStoreFail
	CmdStorePower
	CmdStoreTime
	CmdStoreRate
	CmdStoreScene
	CmdRemoveScene
	CmdAddGroup
	CmdRemoveGroup
	CmdStoreAddress
	CmdStatus
	CmdCheckWork
	CmdCheckLamp
	CmdCheckOperat
	CmdCheckLevel
	CmdCheckAddress
	CmdVersion
	CmdDtr
	CmdType
	CmdPhysical
	CmdCheckFail
	CmdLevel
	CmdMax
	CmdMin
	CmdPower
	CmdFail
	CmdTimeRate
	CmdScene
	CmdGroup0
	CmdGroup8
	CmdRandomh
	CmdRandomm
	CmdRandoml
	CmdCheckReset
)
