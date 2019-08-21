<<<<<<< HEAD
package UHFRFID

/*GetReaderInformation
	The host sends this command to get the reader’s information including reader’s address (Adr), firmware version, reader’s type (Type), supported protocol (Tr_Type), reader power, work frequency, and InventoryScanTime value.
*/
func GetReaderInformation(adr uint8)(Res, Reader){
	send(Req{adr: adr, cmd: getReaderInformation,})
	res := recv()
	const (
		Version = 0
		Type 	= 2
		Tr_Type = 3
		MaxFre 	= 4
		MinFre 	= 5
		Pwr 	= 6
		ScanTime= 7
	)
	reader := Reader{
		Version: 	res.data[Version:Type],
		Type: 		res.data[Type],
		Tr_Type: 	res.data[Tr_Type],
		MaxFre: 	res.data[MaxFre],
		MinFre: 	res.data[MinFre],
		Pwr: 		res.data[Pwr],
		ScanTime: 	res.data[ScanTime],
	}
	return res, reader
}
/*SetRegion
	The host sends this command to change the current region of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetRegion(adr uint8, reader Reader) (Res){
	data := []uint8{reader.MaxFre, reader.MinFre}
	send(Req{adr: adr, cmd: setRegion, data: data,})
	res := recv()
	return res
}
/*SetAddress
	The host sends this command to change the address (Adr) of the reader. 
	The address data is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off. The default value of Adr is 0x00. 
	The range of Adr is 0x00~0xFE. When the host tries to write 0xFF to Adr, the reader will set the value to 0x00 automatically.
*/
func SetAddress(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Adr}
	send(Req{adr: adr, cmd: setAddress, data: data,})
	res := recv()
	return res
}
/*SetScanTime
	The host sends this command to change the value of InventoryScanTime of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetScanTime(adr uint8, reader Reader) (Res){
	data := []uint8{reader.ScanTime}
	send(Req{adr: adr, cmd: setScanTime, data: data,})
	res := recv()
	return res
}
/*SetBaudRate
	The host sends this command to change the value of band rate of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetBaudRate(adr uint8, reader Reader) (Res){
	data := []uint8{reader.BaudRate}
	send(Req{adr: adr, cmd: setBaudRate, data: data,})
	res := recv()
	return res
}
/*SetPower
	The host sends this command to change the power of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetPower(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Pwr}
	send(Req{adr: adr, cmd: setPower, data: data,})
	res := recv()
	return res
}
/*Acousto_opticControl
	The host sends this command to control the LED lights flash and buzzer tweet. 
*/
func Acousto_opticControl(adr uint8, reader Reader) (Res){
	data := []uint8{reader.ActiveT, reader.SilentT, reader.Times}
	send(Req{adr: adr, cmd: acousto_opticControl, data: data,})
	res := recv()
	return res
}
/*SetWiegand
	The host sends this command to change Wiegand parameter of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetWiegand(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Wg_mode, reader.Wg_Data_Inteval, reader.Wg_Pulse_Width, reader.Wg_Pulse_Inteval}
	send(Req{adr: adr, cmd: setWiegand, data: data,})
	res := recv()
	return res
}
/*SetWorkMode
	The host sends this command to set the reader’s in Scan Mode or Trigger Mode. 
	The host can also use this command to define the reader’s output data content and format.
	In Scan Mode or Trigger Mode, the reader can still accept commands from the host. 
	But it will only respond to reader-defined commands. Other commands can not be executed when the reader in Scan Mode or Trigger Mode.
*/
func SetWorkMode(adr uint8, reader Reader) (Res){
	data := reader.WorkMode
	send(Req{adr: adr, cmd: setWorkMode, data: data,})
	res := recv()
	return res
}
/*GetWorkMode
The host sends this command to get the reader’s information including reader’s Wiegand parameter, WorkMode parameter.
*/
func GetWorkMode(adr uint8) (Res, Reader){
	send(Req{adr: adr, cmd: getWorkMode,})
	res := recv()
	const(
		Wg_mode	 		= 0
		Wg_Data_Inteval	= 1
		Wg_Pulse_Width	= 2
		Wg_Pulse_Inteval= 3
		Read_mode		= 4
		Mode_state 		= 5
		Mem_Inven 		= 6
		First_Adr 		= 7
		Word_Num 		= 8
		Tag_Time 		= 9
		Accuracy 		= 10
		OffsetTime 		= 11
	)
	workMode := []uint8{res.data[Read_mode], res.data[Mode_state], res.data[Mem_Inven], res.data[First_Adr], res.data[Word_Num], res.data[Word_Num]}
	reader := Reader{
		Wg_mode: 			res.data[Wg_mode],
		Wg_Data_Inteval: 	res.data[Wg_Data_Inteval],
		Wg_Pulse_Width: 	res.data[Wg_Pulse_Width],
		Wg_Pulse_Inteval: 	res.data[Wg_Pulse_Inteval],
		WorkMode: 		 	workMode,
		Accuracy: 		 	res.data[Accuracy],
		OffsetTime: 		res.data[OffsetTime],
	}
	return res, reader
}
/*SetEasAccuracy
	The host sends this command to set EAS Alarm Accuracy.
*/
func SetEasAccuracy(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Accuracy}
	send(Req{adr: adr, cmd: setEasAccuracy, data: data,})
	res := recv()
	return res
}
/*SyrisResponseOffset
	The host sends this command to set Syris485 response offset time.
*/
func SyrisResponseOffset(adr uint8, reader Reader) (Res){
	data := []uint8{reader.OffsetTime}
	send(Req{adr: adr, cmd: syrisResponseOffset, data: data,})
	res := recv()
	return res
}
/*TriggerOffset
	The host sends this command to set Trigger offset time. 
	This function is only available for reader with firmware version V2.36 and above.
*/
func TriggerOffset(adr uint8, reader Reader) (Res, Reader){
	data := []uint8{reader.TriggerTime}
	send(Req{adr: adr, cmd: triggerOffset, data: data,})
	res := recv()
	const TriggerTime = 0
	reader.TriggerTime = res.data[TriggerTime]
	return res, reader
}
=======
package UHFRFID

/*GetReaderInformation
	The host sends this command to get the reader’s information including reader’s address (Adr), firmware version, reader’s type (Type), supported protocol (Tr_Type), reader power, work frequency, and InventoryScanTime value.
*/
func GetReaderInformation(adr uint8)(Res, Reader){
	send(Req{adr: adr, cmd: getReaderInformation,})
	res := recv()
	const (
		Version = 0
		Type 	= 2
		Tr_Type = 3
		MaxFre 	= 4
		MinFre 	= 5
		Pwr 	= 6
		ScanTime= 7
	)
	reader := Reader{
		Version: 	res.data[Version:Type],
		Type: 		res.data[Type],
		Tr_Type: 	res.data[Tr_Type],
		MaxFre: 	res.data[MaxFre],
		MinFre: 	res.data[MinFre],
		Pwr: 		res.data[Pwr],
		ScanTime: 	res.data[ScanTime],
	}
	return res, reader
}
/*SetRegion
	The host sends this command to change the current region of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetRegion(adr uint8, reader Reader) (Res){
	data := []uint8{reader.MaxFre, reader.MinFre}
	send(Req{adr: adr, cmd: setRegion, data: data,})
	res := recv()
	return res
}
/*SetAddress
	The host sends this command to change the address (Adr) of the reader. 
	The address data is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off. The default value of Adr is 0x00. 
	The range of Adr is 0x00~0xFE. When the host tries to write 0xFF to Adr, the reader will set the value to 0x00 automatically.
*/
func SetAddress(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Adr}
	send(Req{adr: adr, cmd: setAddress, data: data,})
	res := recv()
	return res
}
/*SetScanTime
	The host sends this command to change the value of InventoryScanTime of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetScanTime(adr uint8, reader Reader) (Res){
	data := []uint8{reader.ScanTime}
	send(Req{adr: adr, cmd: setScanTime, data: data,})
	res := recv()
	return res
}
/*SetBaudRate
	The host sends this command to change the value of band rate of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetBaudRate(adr uint8, reader Reader) (Res){
	data := []uint8{reader.BaudRate}
	send(Req{adr: adr, cmd: setBaudRate, data: data,})
	res := recv()
	return res
}
/*SetPower
	The host sends this command to change the power of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetPower(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Pwr}
	send(Req{adr: adr, cmd: setPower, data: data,})
	res := recv()
	return res
}
/*Acousto_opticControl
	The host sends this command to control the LED lights flash and buzzer tweet. 
*/
func Acousto_opticControl(adr uint8, reader Reader) (Res){
	data := []uint8{reader.ActiveT, reader.SilentT, reader.Times}
	send(Req{adr: adr, cmd: acousto_opticControl, data: data,})
	res := recv()
	return res
}
/*SetWiegand
	The host sends this command to change Wiegand parameter of the reader. 
	The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
*/
func SetWiegand(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Wg_mode, reader.Wg_Data_Inteval, reader.Wg_Pulse_Width, reader.Wg_Pulse_Inteval}
	send(Req{adr: adr, cmd: setWiegand, data: data,})
	res := recv()
	return res
}
/*SetWorkMode
	The host sends this command to set the reader’s in Scan Mode or Trigger Mode. 
	The host can also use this command to define the reader’s output data content and format.
	In Scan Mode or Trigger Mode, the reader can still accept commands from the host. 
	But it will only respond to reader-defined commands. Other commands can not be executed when the reader in Scan Mode or Trigger Mode.
*/
func SetWorkMode(adr uint8, reader Reader) (Res){
	data := reader.WorkMode
	send(Req{adr: adr, cmd: setWorkMode, data: data,})
	res := recv()
	return res
}
/*GetWorkMode
The host sends this command to get the reader’s information including reader’s Wiegand parameter, WorkMode parameter.
*/
func GetWorkMode(adr uint8) (Res, Reader){
	send(Req{adr: adr, cmd: getWorkMode,})
	res := recv()
	const(
		Wg_mode	 		= 0
		Wg_Data_Inteval	= 1
		Wg_Pulse_Width	= 2
		Wg_Pulse_Inteval= 3
		Read_mode		= 4
		Mode_state 		= 5
		Mem_Inven 		= 6
		First_Adr 		= 7
		Word_Num 		= 8
		Tag_Time 		= 9
		Accuracy 		= 10
		OffsetTime 		= 11
	)
	workMode := []uint8{res.data[Read_mode], res.data[Mode_state], res.data[Mem_Inven], res.data[First_Adr], res.data[Word_Num], res.data[Word_Num]}
	reader := Reader{
		Wg_mode: 			res.data[Wg_mode],
		Wg_Data_Inteval: 	res.data[Wg_Data_Inteval],
		Wg_Pulse_Width: 	res.data[Wg_Pulse_Width],
		Wg_Pulse_Inteval: 	res.data[Wg_Pulse_Inteval],
		WorkMode: 		 	workMode,
		Accuracy: 		 	res.data[Accuracy],
		OffsetTime: 		res.data[OffsetTime],
	}
	return res, reader
}
/*SetEasAccuracy
	The host sends this command to set EAS Alarm Accuracy.
*/
func SetEasAccuracy(adr uint8, reader Reader) (Res){
	data := []uint8{reader.Accuracy}
	send(Req{adr: adr, cmd: setEasAccuracy, data: data,})
	res := recv()
	return res
}
/*SyrisResponseOffset
	The host sends this command to set Syris485 response offset time.
*/
func SyrisResponseOffset(adr uint8, reader Reader) (Res){
	data := []uint8{reader.OffsetTime}
	send(Req{adr: adr, cmd: syrisResponseOffset, data: data,})
	res := recv()
	return res
}
/*TriggerOffset
	The host sends this command to set Trigger offset time. 
	This function is only available for reader with firmware version V2.36 and above.
*/
func TriggerOffset(adr uint8, reader Reader) (Res, Reader){
	data := []uint8{reader.TriggerTime}
	send(Req{adr: adr, cmd: triggerOffset, data: data,})
	res := recv()
	const TriggerTime = 0
	reader.TriggerTime = res.data[TriggerTime]
	return res, reader
}
>>>>>>> 8/22
