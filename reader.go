package uhf

//GetReaderInformation The host sends this command to get the reader’s information including reader’s address (Adr), firmware version, reader’s type (Type), supported protocol (TrType), reader power, work frequency, and InventoryScanTime value.
func GetReaderInformation(adr uint8) (Res, Reader) {
	send(Req{adr: adr, cmd: CmdGetReaderInformation})
	res := recv()
	const (
		Version  = 0
		Type     = 2
		TrType   = 3
		MaxFre   = 4
		MinFre   = 5
		Pwr      = 6
		ScanTime = 7
	)
	reader := Reader{
		Version:  res.Data[Version:Type],
		Type:     res.Data[Type],
		TrType:   res.Data[TrType],
		MaxFre:   res.Data[MaxFre],
		MinFre:   res.Data[MinFre],
		Pwr:      res.Data[Pwr],
		ScanTime: res.Data[ScanTime],
	}
	return res, reader
}

//SetRegion The host sends this command to change the current region of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
func SetRegion(adr uint8, reader Reader) Res {
	data := []uint8{reader.MaxFre, reader.MinFre}
	send(Req{adr: adr, cmd: CmdSetRegion, data: data})
	res := recv()
	return res
}

//SetAddress The host sends this command to change the address (Adr) of the reader. The address data is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off. The default value of Adr is 0x00. The range of Adr is 0x00~0xFE. When the host tries to write 0xFF to Adr, the reader will set the value to 0x00 automatically.
func SetAddress(adr uint8, reader Reader) Res {
	data := []uint8{reader.Adr}
	send(Req{adr: adr, cmd: CmdSetAddress, data: data})
	res := recv()
	return res
}

//SetScanTime The host sends this command to change the value of InventoryScanTime of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
func SetScanTime(adr uint8, reader Reader) Res {
	data := []uint8{reader.ScanTime}
	send(Req{adr: adr, cmd: CmdSetScanTime, data: data})
	res := recv()
	return res
}

//SetBaudRate The host sends this command to change the value of band rate of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
func SetBaudRate(adr uint8, reader Reader) Res {
	data := []uint8{reader.BaudRate}
	send(Req{adr: adr, cmd: CmdSetBaudRate, data: data})
	res := recv()
	return res
}

//SetPower The host sends this command to change the power of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
func SetPower(adr uint8, reader Reader) Res {
	data := []uint8{reader.Pwr}
	send(Req{adr: adr, cmd: CmdSetPower, data: data})
	res := recv()
	return res
}

//AcoustoOpticControl  The host sends this command to control the LED lights flash and buzzer tweet.
func AcoustoOpticControl(adr uint8, reader Reader) Res {
	data := []uint8{reader.ActiveT, reader.SilentT, reader.Times}
	send(Req{adr: adr, cmd: CmdAcoustoOpticControl, data: data})
	res := recv()
	return res
}

//SetWiegand The host sends this command to change Wiegand parameter of the reader. The value is stored in the reader’s inner EEPROM and is nonvolatile after reader powered off.
func SetWiegand(adr uint8, reader Reader) Res {
	data := []uint8{reader.WgMode, reader.WgDataInteval, reader.WgPulseWidth, reader.WgPulseInteval}
	send(Req{adr: adr, cmd: CmdSetWiegand, data: data})
	res := recv()
	return res
}

//SetWorkMode The host sends this command to set the reader’s in Scan Mode or Trigger Mode. The host can also use this command to define the reader’s output data content and format. In Scan Mode or Trigger Mode, the reader can still accept commands from the host. But it will only respond to reader-defined commands. Other commands can not be executed when the reader in Scan Mode or Trigger Mode.
func SetWorkMode(adr uint8, reader Reader) Res {
	data := reader.WorkMode
	send(Req{adr: adr, cmd: CmdSetWorkMode, data: data})
	res := recv()
	return res
}

//GetWorkMode The host sends this command to get the reader’s information including reader’s Wiegand parameter, WorkMode parameter.
func GetWorkMode(adr uint8) (Res, Reader) {
	send(Req{adr: adr, cmd: CmdGetWorkMode})
	res := recv()
	const (
		WgMode         = 0
		WgDataInteval  = 1
		WgPulseWidth   = 2
		WgPulseInteval = 3
		ReadMode       = 4
		ModeState      = 5
		MemInven       = 6
		FirstAdr       = 7
		WordNum        = 8
		TagTime        = 9
		Accuracy       = 10
		OffsetTime     = 11
	)
	workMode := []uint8{res.Data[ReadMode], res.Data[ModeState], res.Data[MemInven], res.Data[FirstAdr], res.Data[WordNum], res.Data[WordNum]}
	reader := Reader{
		WgMode:         res.Data[WgMode],
		WgDataInteval:  res.Data[WgDataInteval],
		WgPulseWidth:   res.Data[WgPulseWidth],
		WgPulseInteval: res.Data[WgPulseInteval],
		WorkMode:       workMode,
		Accuracy:       res.Data[Accuracy],
		OffsetTime:     res.Data[OffsetTime],
	}
	return res, reader
}

//SetEasAccuracy The host sends this command to set EAS Alarm Accuracy.
func SetEasAccuracy(adr uint8, reader Reader) Res {
	data := []uint8{reader.Accuracy}
	send(Req{adr: adr, cmd: CmdSetEasAccuracy, data: data})
	res := recv()
	return res
}

//SyrisResponseOffset The host sends this command to set Syris485 response offset time.
func SyrisResponseOffset(adr uint8, reader Reader) Res {
	data := []uint8{reader.OffsetTime}
	send(Req{adr: adr, cmd: CmdSyrisResponseOffset, data: data})
	res := recv()
	return res
}

//TriggerOffset The host sends this command to set Trigger offset time. This function is only available for reader with firmware version V2.36 and above.
func TriggerOffset(adr uint8, reader Reader) (Res, Reader) {
	data := []uint8{reader.TriggerTime}
	send(Req{adr: adr, cmd: CmdTriggerOffset, data: data})
	res := recv()
	const TriggerTime = 0
	reader.TriggerTime = res.Data[TriggerTime]
	return res, reader
}
