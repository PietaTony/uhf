package uhf

//Req request
type Req struct {
	len  uint8
	adr  uint8
	cmd  uint8
	data []uint8
	lsb  uint8
	msb  uint8
}

func (m *Req) getBytesWithoutCRC() []uint8 {
	m.len = (uint8)(cmdMinSize + len(m.data))
	slice := []uint8{m.len, m.adr, m.cmd}
	return append(slice, m.data...)
}

//GetBytes get request by bytes
func (m *Req) GetBytes() []uint8 {
	slice := m.getBytesWithoutCRC()
	m.msb, m.lsb = getCRC16(slice)
	return append(slice, m.lsb, m.msb)
}

//GetString get request by string
func (m *Req) GetString() string {
	s := "Request:" + "\n"
	s += "\tAdr:\t" + string(GetStr(m.adr)) + "\n"
	s += "\tCmd:\t" + getCmdStr(m.cmd) + "\n"
	if len(m.data) > 0 {
		s += "\tData:\t" + string(GetStr(m.data)) + "\n"
	}
	s += "\tAll:\t" + GetStr(m.GetBytes()) + "\n"
	return s
}
