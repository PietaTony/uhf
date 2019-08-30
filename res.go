package reader

import "errors"

//Res response
type Res struct {
	Len    uint8
	Adr    uint8
	Cmd    uint8
	Status uint8
	Data   []uint8
	Lsb    uint8
	Msb    uint8
}

func (m *Res) initRes(buf []uint8) error {
	const (
		lenPos    = 0
		adrPos    = 1
		cmdPos    = 2
		statusPos = 3
		dataPos   = 4
		crcSize   = 2
		lsbPos    = 2
		msbPos    = 1
	)
	n := len(buf)
	m.Len = buf[lenPos]
	m.Adr = buf[adrPos]
	m.Cmd = buf[cmdPos]
	m.Status = buf[statusPos]
	if n > recmdMinSize+1 {
		m.Data = buf[dataPos : n-crcSize]
	}
	m.Lsb = buf[n-lsbPos]
	m.Msb = buf[n-msbPos]
	MSB, LSB := getCRC16(m.getBytesWithoutCRC())
	if MSB == m.Msb && LSB == m.Lsb {
		return nil
	}
	return errors.New("CRC Error")

}

func (m *Res) getBytesWithoutCRC() []uint8 {
	slice := []uint8{m.Len, m.Adr, m.Cmd, m.Status}
	return append(slice, m.Data...)
}

//GetBytes get response by bytes
func (m *Res) GetBytes() []uint8 {
	slice := m.getBytesWithoutCRC()
	m.Msb, m.Lsb = getCRC16(slice)
	return append(slice, m.Lsb, m.Msb)
}

//GetString get response by string
func (m *Res) GetString() string {
	s := "Response:" + "\n"
	s += "\tAdr:\t" + string(GetStr(m.Adr)) + "\n"
	s += "\tCmd:\t" + getCmdStr(m.Cmd) + "\n"
	s += "\tStatus:\t" + getStatusStr(m.Status) + "\n"
	if len(m.Data) > 0 {
		s += "\tData:\t" + string(GetStr(m.Data)) + "\n"
	}
	s += "\tAll:\t" + GetStr(m.GetBytes()) + "\n"
	return s
}
