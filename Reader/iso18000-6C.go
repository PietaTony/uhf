package UHFRFID

/*Inventory
The command function is used to inventory tags in the effective field and get their EPC or TID values. The reader executes an Inventory command and gets tag’s EPC before any other operation.
The user may accord need to establish this command the first biggest running time (Inventory scan time), before the command enquires. The reader completes command execution in inventory ScanTime (not including host sending data time) except inventory command after receiving host command and returns the results.
The default value is 0x0A (corresponding to 10*100ms=1s).
The value range is 0x03~0xFF (corresponding to 3*100ms~255*100ms). In various environments, the actual inventory scan time may be 0~75ms longer than the InventoryScanTime defined.
If the inventory scan time establishes excessively short, possibly will inventory no tag appear in inventory scan time.
*/
func Inventory(adr uint8, spec Spec, TID Memory) (Res, []Memory) {
	data := []uint8{spec.Adr, TID.Len}
	send(Req{adr: adr, cmd: inventory, data: data})
	res := recv()
	if res.Len == recmdMinSize {
		var EPSs []Memory
		return res, EPSs
	}

	const (
		Num    = 0
		EPCPos = 1
	)
	EPCSize := int(res.Data[Num])
	EPCs := make([]Memory, EPCSize)

	n := EPCPos
	for i := 0; i < int(EPCSize); i++ {
		EPCs[i].Len = res.Data[n]
		n++
		EPCs[i].Data = res.Data[n : n+int(EPCs[i].Len)]
		n += int(EPCs[i].Len)
	}

	return res, EPCs
}
func InventoryAll(adr uint8) (Res, []Memory) {
	data := []uint8{}
	send(Req{adr: adr, cmd: inventory, data: data})
	res := recv()
	if res.Len == recmdMinSize {
		var EPSs []Memory
		return res, EPSs
	}

	const (
		Num    = 0
		EPCPos = 1
	)
	EPCSize := int(res.Data[Num])
	EPCs := make([]Memory, EPCSize)

	n := EPCPos
	for i := 0; i < int(EPCSize); i++ {
		EPCs[i].Len = res.Data[n]
		n++
		EPCs[i].Data = res.Data[n : n+int(EPCs[i].Len)]
		n += int(EPCs[i].Len)
	}

	return res, EPCs
}

/*ReadData
The command is used to read part or all of a Tag’s Password, EPC, TID, or User memory.
To the word as a unit, start to read data from the designated address.
*/
func ReadData(adr uint8, tag Tag, spec Spec, mask Mask) (Res, []uint8) {
	data := []uint8{tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, spec.Name, spec.Adr, spec.Mem.Len)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: readData, data: data})
	res := recv()
	return res, res.Data
}

/*WriteData
The command is used to write several words in a Tag’s Reserved, EPC, TID, or User memory.
*/
func WriteData(adr uint8, tag Tag, spec Spec, mask Mask) Res {
	data := []uint8{spec.Mem.Len, tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, spec.Name, spec.WordPtr)
	data = append(data, spec.Mem.Data...)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: writeData, data: data})
	res := recv()
	return res
}

/*WriteEPC
The command is used to write EPC number in a Tag’s EPC memory.
Random write one tag in the effective field.
*/
func WriteEPC(adr uint8, pwd []uint8, spec Spec) Res {
	data := []uint8{spec.Mem.Len}
	data = append(data, pwd...)
	data = append(data, spec.Mem.Data...)
	send(Req{adr: adr, cmd: writeEPC, data: data})
	res := recv()
	return res
}

/*KillTag
The command is used to kill tag.
After the tag killed, it never process command.
*/
func KillTag(adr uint8, tag Tag, mask Mask) Res {
	data := []uint8{tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: killTag, data: data})
	res := recv()
	return res
}

/*Lock
The Lock command Lock reversibly or permanently locks a password or an entire EPC, TID, or User memory bank in a readable/writeable or unreadable/unwriteable state.
Once tag’s password memory establishes to forever may be readable and writable or unreadable and unwriteable, then later cannot change its read-write protection again.
Tag’s EPC memory, TID memory or user memory, if establishes to forever may be writeable or unwriteable, then later cannot change its read-write	protection again.
If sends the command to want forcefully to change the above several states, then the tag will return to the error code.
When the tag’s memory established in a readable/writeable state, the command must give the Access
Password, so tag’s Access Password is not zero.
*/
func Lock(adr uint8, tag Tag, prt Protect, mask Mask) Res {
	data := []uint8{tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, prt.Select, prt.SetProtect)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: lock, data: data})
	res := recv()
	return res
}

/*BlockErase
The command is used to erase multiple words in a Tag’s Password, EPC, TID, or User memory.
*/
func BlockErase(adr uint8, tag Tag, spec Spec, mask Mask) Res {
	data := []uint8{tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, spec.Name, spec.WordPtr, spec.Mem.Len)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: blockErase, data: data})
	res := recv()
	return res
}

/*ReadProtect
The command is used to set designated tag read protection.
After the tag protected, it never process command.
Even if inventory tag, reader can not get the EPC number.
The read protection can be removed by executing Reset
ReadProtect. Only NXP's UCODE EPC G2X tags valid.
*/
func ReadProtect(adr uint8, tag Tag, mask Mask) Res {
	data := []uint8{tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: readProtect, data: data})
	res := recv()
	return res
}

/*ReadProtect Without EPC
The command is used to random set random one tag read protection in the effective field.
The tag must be having the same access password.
Only NXP's UCODE EPC G2X tags valid.
*/
func ReadProtectWithoutEPC(adr uint8, pwd []uint8) Res {
	data := pwd
	send(Req{adr: adr, cmd: readProtectWithoutEPC, data: data})
	res := recv()
	return res
}

/*ResetReadProtect
The command is used to remove only one tag read protection in the effective field.
The tag must be having the same access password.
Only NXP's UCODE EPC G2X tags valid.
*/
func ResetReadProtect(adr uint8, pwd []uint8) Res {
	data := pwd
	send(Req{adr: adr, cmd: resetReadProtect, data: data})
	res := recv()
	return res
}

/*CheckReadProtect
The command is used to check only one tag in the effective field, whether the tag is protected.
It can not check the tag whether the tag support protection setting. Only NXP's UCODE EPC G2X tags valid.
*/
func CheckReadProtect(adr uint8) (Res, bool) {
	send(Req{adr: adr, cmd: checkReadProtect})
	res := recv()
	const (
		ReadProtect = 0
		unprotected = 1
	)
	return res, res.Data[ReadProtect] == unprotected
}

/*EASAlarm
The function is used to set or reset the EAS status bit of designated tag.
Only NXP's UCODE EPC G2X tags valid.
*/
func EASAlarm(adr uint8) Res {
	send(Req{adr: adr, cmd: _EASAlarm})
	res := recv()
	return res
}

/*CheckEASAlarm
The function is used to check EAS status bit of any tag in the effective field.
Only NXP's UCODE EPC G2X tags valid.
*/
func CheckEASAlarm(adr uint8) Res {
	send(Req{adr: adr, cmd: checkEASAlarm})
	res := recv()
	return res
}

/*UserBlockLock
The command is used to permanently lock the designated data in designated tag’s user memory.
Block Lock command supports an additional locking mechanism, which allows the locking of individual 32 bit blocks (rows) in the 224 bit User Memory.
Once locked these locks cannot be unlocked.
Only NXP's UCODE EPC G2X tags valid.
*/
func UserBlockLock(adr uint8, tag Tag, spec Spec, mask Mask) Res {
	data := []uint8{tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, tag.Pwd.Data...)
	data = append(data, spec.WordPtr) //Each EEPROM row can be addressed by either of the two related WordPointers: Either of two WordPointers can address one single User Memory row
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: userBlockLock, data: data})
	res := recv()
	return res
}

/*InventorySingle
 */
func InventorySingle(adr uint8) (Res, uint8, Memory) {
	send(Req{adr: adr, cmd: inventorySingle})
	res := recv()
	const (
		Num    = 0
		EPCPos = 1
	)
	EPC := Memory{
		Len:  uint8(len(res.Data[EPCPos:])),
		Data: res.Data[EPCPos:],
	}
	return res, res.Data[Num], EPC
}

/*BlockWrite
The command is used to write multiple words in a Tag’s Reserved, EPC, TID, or User memory.
*/
func BlockWrite(adr uint8, tag Tag, spec Spec, mask Mask) Res {
	data := []uint8{spec.Mem.Len, tag.EPC.Len}
	data = append(data, tag.EPC.Data...)
	data = append(data, spec.Name, spec.WordPtr)
	data = append(data, spec.Mem.Data...)
	data = append(data, tag.Pwd.Data...)
	data = append(data, mask.Adr, mask.Len)
	send(Req{adr: adr, cmd: blockWrite, data: data})
	res := recv()
	return res
}
