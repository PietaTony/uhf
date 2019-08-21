package UHFRFID

const recmdMinSize = 5
const (
	NoTagOperable 							=	0xFB
)
func getStatusStr(status uint8) (string) {
	switch status{
	//Res Status
	case 0x00:	return "Success"
	case 0x01:	return "Return Before Inventory Finished"
	case 0x02:	return "The Inventory-scan-time Overflow"
	case 0x03:	return "More Data"
	case 0x04:	return "Reader Module Flash Is Full"
	case 0x05:	return "Access Password Error"
	case 0x09:	return "Kill Tag Error"
	case 0x0A:	return "Kill Password Error Can't Be Zero"
	case 0x0B:	return "Tag Not Support The Command"
	case 0x0C:	return "Use The Command, Access Password Can't Be Zero"
	case 0x0D:	return "Tag Is Protected, Cannot Set It Again"
	case 0x0E:	return "Tag Is Unprotected, No Need To Reset It"
	case 0x10:	return "There Is Some Locked Bytes, Write Fail"
	case 0x11:	return "Can Not Lock It"
	case 0x12:	return "Be Locked, Cannot Lock It Again"
	case 0x13:	return "Save Fail, Can Use Before Power"
	case 0x14:	return "Cannot Adjust"
	case 0x15:	return "Return Before Inventory Finished"
	case 0x16:	return "Inventory-Scan-Time Overflow"
	case 0x17:	return "More Data"
	case 0x18:	return "Reader Module Flash Is Full"
	case 0x19:	return "Not Support Command Or Access Password"
	case 0xF9:	return "Command Execute Error"
	case 0xFA:	return "Get Tag, Poor Communication, Inoperable"
	case 0xFB:	return "No Tag Operable"
	case 0xFC:	return "Tag Return Error Code"
	case 0xFD:	return "Command Length Wrong"
	case 0xFE:	return "Illegal Command"
	case 0xFF:	return "Parameter Error"
	default:	return "Status Not Found"
	}
}
