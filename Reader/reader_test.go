package UHFRFID

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestInventoryAll(t *testing.T) {
	Begin("/dev/tty.SLAB_USBtoUART", 57600)
	defer Close()

	adr := uint8(0x00)
	/*
		EPCData := []uint8{0xE2, 0x00, 0x00, 0x1B, 0x41, 0x15, 0x02, 0x23, 0x16, 0x20, 0xC9, 0x4B}
		EPC := Memory{Len: (uint8)(len(EPCData)/2), Data: EPCData,}
		pwd := Memory{Data: []uint8{0x00, 0x00},}
		tag := Tag{EPC: EPC, Pwd: pwd,}

		specData := Memory{Len: 0x04,}
		spec := Spec{Name: UserMem, Adr: 0x00, Mem: specData,}

		mask := Mask{Adr: 0x0 0, Len: 0x00,}

		//PrintBytes(ReadData(adr, tag, spec, mask))
	*/
	for i := 0; i < 5; i++ {
		res, _ := InventoryAll(adr)
		a, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(string(a))
		time.Sleep(1000 * time.Millisecond)
	}
}
