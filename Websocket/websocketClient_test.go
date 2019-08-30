package protocal

import (
	"log"
	"testing"

	"github.com/gorilla/websocket"
)

func TestClient(t *testing.T) {

	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8899/echo", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hello ithome30day"))
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
}
