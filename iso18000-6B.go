<<<<<<< HEAD
package UHFRFID

/*InventorySignal6B
	The command is used to Inventory only one tag in the effective field and get their ID values. 
	If more than one tag in the effective field at the same time, reader may be get nothing.
*/
func InventorySignal6B(adr uint8) (Res, UID){
	send(Req{adr: adr, cmd: inventorySignal6B,})
	res := recv()
	return res, res.data
}
/*InventoryMultiple6B
	The command is used to according to the given conditions Inventory tags in the effective field and get their ID values
	Parameter Connect:
		Condition: The condition of detecting tags.
			0x00: equal condition.
			0x01: unequal condition.
			0x02: greater than condition.
			0x03: lower than condition.
		Address: The tag’s start address to compare.
		Mask: It pointed to the data is used to compare. Highest bit in the mask correspond with the far-left byte in the Condition Content. The corresponding bit in the mask is 1 to compare the bit in the Condition Content with the corresponding byte in the tag. The corresponding bit in the mask is 0, not compare.
		Word_data: 8 bytes. It pointed to the array is used to compare.
*/
func InventoryMultiple6B(adr uint8, condition uint8, dataAdr uint8, mask uint8, wordData []uint8) (Res, []UID){
	data := []uint8{condition, dataAdr, mask}
	data = append(data, wordData...)
	send(Req{adr: adr, cmd: inventoryMultiple6B, data: data,})
	res := recv()
	const(
		Num 	= 0
		UIDPos 	= 1
		UIDSize = 8
	)
	var UIDs []UID
	for i := uint8(0); i<data[Num]; i++{
		UIDs[i] = data[UIDPos: UIDPos + UIDSize*(i + 1)]
	}
	return res, UIDs
}
/*ReadData6B
	The command is used to start to read several bytes from the designated address.
	Parameter Connect:
		Address: The tag’s start byte address to read. The range is 0~223. Otherwise, it returns the parameters error message.
		Num: In byte units. It specifies the number of 8-bit bytes to be read. The value range is 1~32, and Address + Num must be less than 224. Otherwise, it returns the parameters error message.
		ID: 8 bytes, it is 6B tag’s UID. The low byte is fist
*/
func ReadData6B(adr uint8, dataAdr uint8, ID UID, num uint8) (Res, []uint8){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	data = append(data, num)
	send(Req{adr: adr, cmd: readData6B, data: data,})
	res := recv()
	return res, res.data
}
/*WriteData6B
	The command is used to start to write several bytes from the designated address.
	Parameter Connect:
		Address: The tag’s start byte address to write. The range is 8~223. Otherwise, it returns the parameters error message.
		ID: 8 bytes, it is 6B tag’s UID. The low byte is fist.
		Wdata: It pointed to the array to write, range is 1~32. If Address + WriteDataLen greater than 224, or Wdata greater than 32 or is zero, reader will return parameter error message. The high bytes of Wdata write in the low address in tag.
*/
func WriteData6B(adr uint8, dataAdr uint8, ID UID, wData []uint8) (Res, []uint8){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	data = append(data, wData...)
	send(Req{adr: adr, cmd: writeData6B, data: data,})
	res := recv()
	return res, res.data
}
/*CheckLock6B
	The command is used to check whether the designated byte is locked.
	Parameter Connect:
		Address: The tag’s byte address to check lock. The range is 0~223. Otherwise, it returns the parameters error message.
		ID: 8 bytes, it is 6B tag’s UID. The low byte is fist.
*/
func CheckLock6B(adr uint8, dataAdr uint8, ID UID) (Res, bool){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	send(Req{adr: adr, cmd: checkLock6B, data: data,})
	res := recv()
	const (
		LockState = 0
		Locked 	  = 1
	)
	return res, (res.data[LockState] == Locked)
}
/*Lock6B
	The command is used to lock the designated byte.
*/
func Lock6B(adr uint8, dataAdr uint8, ID UID) (Res){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	send(Req{adr: adr, cmd: lock6B, data: data,})
	res := recv()
	return res
}
=======
package UHFRFID

/*InventorySignal6B
	The command is used to Inventory only one tag in the effective field and get their ID values. 
	If more than one tag in the effective field at the same time, reader may be get nothing.
*/
func InventorySignal6B(adr uint8) (Res, UID){
	send(Req{adr: adr, cmd: inventorySignal6B,})
	res := recv()
	return res, res.data
}
/*InventoryMultiple6B
	The command is used to according to the given conditions Inventory tags in the effective field and get their ID values
	Parameter Connect:
		Condition: The condition of detecting tags.
			0x00: equal condition.
			0x01: unequal condition.
			0x02: greater than condition.
			0x03: lower than condition.
		Address: The tag’s start address to compare.
		Mask: It pointed to the data is used to compare. Highest bit in the mask correspond with the far-left byte in the Condition Content. The corresponding bit in the mask is 1 to compare the bit in the Condition Content with the corresponding byte in the tag. The corresponding bit in the mask is 0, not compare.
		Word_data: 8 bytes. It pointed to the array is used to compare.
*/
func InventoryMultiple6B(adr uint8, condition uint8, dataAdr uint8, mask uint8, wordData []uint8) (Res, []UID){
	data := []uint8{condition, dataAdr, mask}
	data = append(data, wordData...)
	send(Req{adr: adr, cmd: inventoryMultiple6B, data: data,})
	res := recv()
	const(
		Num 	= 0
		UIDPos 	= 1
		UIDSize = 8
	)
	var UIDs []UID
	for i := uint8(0); i<data[Num]; i++{
		UIDs[i] = data[UIDPos: UIDPos + UIDSize*(i + 1)]
	}
	return res, UIDs
}
/*ReadData6B
	The command is used to start to read several bytes from the designated address.
	Parameter Connect:
		Address: The tag’s start byte address to read. The range is 0~223. Otherwise, it returns the parameters error message.
		Num: In byte units. It specifies the number of 8-bit bytes to be read. The value range is 1~32, and Address + Num must be less than 224. Otherwise, it returns the parameters error message.
		ID: 8 bytes, it is 6B tag’s UID. The low byte is fist
*/
func ReadData6B(adr uint8, dataAdr uint8, ID UID, num uint8) (Res, []uint8){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	data = append(data, num)
	send(Req{adr: adr, cmd: readData6B, data: data,})
	res := recv()
	return res, res.data
}
/*WriteData6B
	The command is used to start to write several bytes from the designated address.
	Parameter Connect:
		Address: The tag’s start byte address to write. The range is 8~223. Otherwise, it returns the parameters error message.
		ID: 8 bytes, it is 6B tag’s UID. The low byte is fist.
		Wdata: It pointed to the array to write, range is 1~32. If Address + WriteDataLen greater than 224, or Wdata greater than 32 or is zero, reader will return parameter error message. The high bytes of Wdata write in the low address in tag.
*/
func WriteData6B(adr uint8, dataAdr uint8, ID UID, wData []uint8) (Res, []uint8){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	data = append(data, wData...)
	send(Req{adr: adr, cmd: writeData6B, data: data,})
	res := recv()
	return res, res.data
}
/*CheckLock6B
	The command is used to check whether the designated byte is locked.
	Parameter Connect:
		Address: The tag’s byte address to check lock. The range is 0~223. Otherwise, it returns the parameters error message.
		ID: 8 bytes, it is 6B tag’s UID. The low byte is fist.
*/
func CheckLock6B(adr uint8, dataAdr uint8, ID UID) (Res, bool){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	send(Req{adr: adr, cmd: checkLock6B, data: data,})
	res := recv()
	const (
		LockState = 0
		Locked 	  = 1
	)
	return res, (res.data[LockState] == Locked)
}
/*Lock6B
	The command is used to lock the designated byte.
*/
func Lock6B(adr uint8, dataAdr uint8, ID UID) (Res){
	data := []uint8{dataAdr}
	data = append(data, ID...)
	send(Req{adr: adr, cmd: lock6B, data: data,})
	res := recv()
	return res
}
>>>>>>> 8/22
