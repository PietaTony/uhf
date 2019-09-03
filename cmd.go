package uhf

const cmdMinSize = 4

//CMD
const (
	//ISO18000-6C CMD
	CmdInventory             = 0x01
	CmdReadData              = 0x02
	CmdWriteData             = 0x03
	CmdWriteEPC              = 0x04
	CmdKillTag               = 0x05
	CmdLock                  = 0x06
	CmdBlockErase            = 0x07
	CmdReadProtect           = 0x08
	CmdReadProtectWithoutEPC = 0x09
	CmdResetReadProtect      = 0x0A
	CmdCheckReadProtect      = 0x0B
	CmdEASAlarm              = 0x0C
	CmdCheckEASAlarm         = 0x0D
	CmdUserBlockLock         = 0x0E
	CmdInventorySingle       = 0x0F
	CmdBlockWrite            = 0x10
	//ISO18000-6B CMD
	CmdInventorySignal6B   = 0x50
	CmdInventoryMultiple6B = 0x51
	CmdReadData6B          = 0x52
	CmdWriteData6B         = 0x53
	CmdCheckLock6B         = 0x54
	CmdLock6B              = 0x55
	//Reader Defined CMD
	CmdGetReaderInformation = 0x21
	CmdSetRegion            = 0x22
	CmdSetAddress           = 0x24
	CmdSetScanTime          = 0x25
	CmdSetBaudRate          = 0x28
	CmdSetPower             = 0x2F
	CmdAcoustoOpticControl  = 0x33
	CmdSetWiegand           = 0x34
	CmdSetWorkMode          = 0x35
	CmdGetWorkMode          = 0x36
	CmdSetEasAccuracy       = 0x37
	CmdSyrisResponseOffset  = 0x38
	CmdTriggerOffset        = 0x3B
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
