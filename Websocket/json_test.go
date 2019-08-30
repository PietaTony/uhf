package protocal

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	ID             int
	Name           string
	Money          float64
	Skills         []string
	Relationship   map[string]string
	Identification Identification
}

type Identification struct {
	Phone bool
	Email bool
}

func TestJson(t *testing.T) {
	user := User{
		ID:     1,
		Name:   "Tony",
		Skills: []string{"program", "rich", "play"},
		Relationship: map[string]string{
			"Dad": "Hulk",
			"Mon": "Natasha",
		},
		Identification: Identification{
			Phone: true,
			Email: false,
		},
	}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
