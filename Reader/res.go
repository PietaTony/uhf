package UHFRFID

import "errors"

type Res struct{
	len 	uint8
	adr		uint8
	cmd 	uint8
	status 	uint8
	data 	[]uint8
	lsb	 	uint8
	msb	 	uint8
}

func (m *Res)initRes(buf []uint8)(error){
	const(
		lenPos 		= 0
		adrPos 		= 1
		cmdPos 		= 2
		statusPos 	= 3
		dataPos 	= 4
		crcSize 	= 2
		lsbPos 		= 2
		msbPos 		= 1
	)
	n := len(buf)
	m.len = buf[lenPos]
	m.adr = buf[adrPos]
	m.cmd = buf[cmdPos]
	m.status = buf[statusPos]
	if n > recmdMinSize+1 {
		m.data = buf[dataPos: n-crcSize]
	}
	m.lsb = buf[n-lsbPos]
	m.msb = buf[n-msbPos]
	MSB, LSB := getCRC16(m.getBytesWithoutCRC())
	if MSB==m.msb && LSB==m.lsb{
		return nil
	} else {
		return errors.New("CRC Error")
	}
}

func (m *Res)getBytesWithoutCRC() ([]uint8){
	slice := []uint8{m.len, m.adr, m.cmd, m.status}
	return append(slice, m.data...)
}

func (m *Res)GetBytes() ([]uint8){
	slice := m.getBytesWithoutCRC()
	m.msb, m.lsb = getCRC16(slice)
	return append(slice, m.lsb, m.msb)
}

func (m *Res)GetString()(string){
	s := "Response:" + "\n"
	s += "\tAdr:\t" + string(GetStr(m.adr)) + "\n"
	s += "\tCmd:\t" + getCmdStr(m.cmd) + "\n"
	s += "\tStatus:\t" + getStatusStr(m.status) + "\n"
	if len(m.data) > 0{
		s += "\tData:\t" + string(GetStr(m.data)) + "\n"
	}
	s += "\tAll:\t" + GetStr(m.GetBytes()) + "\n"
	return s
}