package protocal

import (
	"encoding/json"
	"log"
	"testing"
	"time"
	UHFRFID "uhf-rfid/Reader"

	"github.com/gorilla/websocket"
)

func TestClient(t *testing.T) {
	UHFRFID.Begin("/dev/tty.SLAB_USBtoUART", 57600)
	defer UHFRFID.Close()
	adr := uint8(0x00)

	res, _ := UHFRFID.InventoryAll(adr)
	data, err := json.Marshal(res)
	if err != nil {
		log.Println("error:", err)
	}

	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8899/echo", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println(err)
		return
	}
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("receive: %s\n", msg)
	time.Sleep(1000 * time.Millisecond)
}
