package uhf

import (
	"fmt"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
	"github.com/sigurn/crc16"
)

//Options Reader's serial options
var Options serial.OpenOptions

//Port Reader serial port
var Port io.ReadWriteCloser

const (
	dataBits        = 8
	stopBits        = 1
	minimumReadSize = 4
)

//Begin begin serial port.
func Begin(portName string, baudRate uint) {
	Options = serial.OpenOptions{
		PortName:        portName,
		BaudRate:        baudRate,
		DataBits:        dataBits,
		StopBits:        stopBits,
		MinimumReadSize: minimumReadSize,
	}
	var err error
	Port, err = serial.Open(Options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
}

//Close close serial port
func Close() {
	Port.Close()
}

func send(req Req) {
	b := req.GetBytes()
	_, err := Port.Write(b)

	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println(req.GetString())
}

func recv() Res {
	buf := make([]byte, 254)
	n, err := Port.Read(buf)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	var res Res
	err = res.initRes(buf[:n])
	if err != nil {
		log.Fatalf("res.initRes: %v", err)
	}

	fmt.Println(res.GetString())
	return res
}

//return MSB LSB
func getCRC16(data []uint8) (uint8, uint8) {
	crc := crc16.Checksum(data, crc16.MakeTable(crc16.CRC16_MCRF4XX))
	return uint8(crc >> 8), uint8(crc & 0xff)
}

func getByte(data uint8) string {
	if data < 16 {
		return fmt.Sprintf("0x0%X ", data)
	}
	return fmt.Sprintf("0x%X ", data)

}

func getBytes(data []uint8) string {
	var s string
	for _, v := range data {
		s += getByte(v)
	}
	return s
}

//GetStr get byte or bytes string
func GetStr(data interface{}) string {
	switch v := data.(type) {
	case uint8:
		return getByte(v)
	case []uint8:
		return getBytes(v)
	default:
		return ""
	}
}

//PrintBytes print bytes data
func PrintBytes(data []uint8) {
	fmt.Print(GetStr(data))
}
