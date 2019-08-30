package uhf

const cmdMinSize = 4

const (
	//ISO18000-6C CMD
	inventory             = 0x01
	readData              = 0x02
	writeData             = 0x03
	writeEPC              = 0x04
	killTag               = 0x05
	lock                  = 0x06
	blockErase            = 0x07
	readProtect           = 0x08
	readProtectWithoutEPC = 0x09
	resetReadProtect      = 0x0A
	checkReadProtect      = 0x0B
	_EASAlarm             = 0x0C
	checkEASAlarm         = 0x0D
	userBlockLock         = 0x0E
	inventorySingle       = 0x0F
	blockWrite            = 0x10
	//ISO18000-6B CMD
	inventorySignal6B   = 0x50
	inventoryMultiple6B = 0x51
	readData6B          = 0x52
	writeData6B         = 0x53
	checkLock6B         = 0x54
	lock6B              = 0x55
	//Reader Defined CMD
	getReaderInformation = 0x21
	setRegion            = 0x22
	setAddress           = 0x24
	setScanTime          = 0x25
	setBaudRate          = 0x28
	setPower             = 0x2F
	acoustoOpticControl  = 0x33
	setWiegand           = 0x34
	setWorkMode          = 0x35
	getWorkMode          = 0x36
	setEasAccuracy       = 0x37
	syrisResponseOffset  = 0x38
	triggerOffset        = 0x3B
)

func getCmdStr(cmd uint8) string {
	switch cmd {
	//ISO18000-6C CMD
	case 0x01:
		return "Inventory"
	case 0x02:
		return "Read Data"
	case 0x03:
		return "Write Data"
	case 0x04:
		return "Write EPC"
	case 0x05:
		return "Kill Tag"
	case 0x06:
		return "Lock"
	case 0x07:
		return "Block Erase"
	case 0x08:
		return "Read Protect"
	case 0x09:
		return "Read Protect Without EPC"
	case 0x0A:
		return "Reset Read Protect"
	case 0x0B:
		return "Check Read Protect"
	case 0x0C:
		return "EAS Alarm"
	case 0x0D:
		return "Check EAS Alarm"
	case 0x0E:
		return "User Block Lock"
	case 0x0F:
		return "Inventory Single"
	case 0x10:
		return "Block Write"
		//ISO18000-6B CMD
	case 0x50:
		return "Inventory Signal 6B"
	case 0x51:
		return "Inventory Multiple 6B"
	case 0x52:
		return "Read Data 6B"
	case 0x53:
		return "Write Data 6B"
	case 0x54:
		return "Check Lock 6B"
	case 0x55:
		return "Lock 6B"
		//Reader Defined CMD
	case 0x21:
		return "Get Reader Information"
	case 0x22:
		return "Set Region"
	case 0x24:
		return "Set Address"
	case 0x25:
		return "Set ScanTime"
	case 0x28:
		return "Set BaudRate"
	case 0x2F:
		return "Set Power"
	case 0x33:
		return "Acousto-optic Control"
	case 0x34:
		return "Set Wiegand"
	case 0x35:
		return "Set WorkMode"
	case 0x36:
		return "Get WorkMode"
	case 0x37:
		return "Set Eas Accuracy"
	case 0x38:
		return "Syris Response Offset"
	case 0x3B:
		return "Trigger Offset"
	default:
		return "Cmd Not Found"
	}
}
