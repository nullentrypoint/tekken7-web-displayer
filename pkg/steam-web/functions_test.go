package steamweb

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testItem struct {
	steamID  string
	nickName string
}

const (
	exmapleSteamID  = "76561197980983067"
	exampleNickName = "HITMAN"
)

func TestParse(t *testing.T) {
	items := []*testItem{
		{
			steamID:  exmapleSteamID,
			nickName: exampleNickName,
		},
		{
			steamID:  "76561199002348717",
			nickName: "Che ☭ Guevara",
		},
		{
			steamID:  "76561198064687587",
			nickName: "patriarch",
		},
	}

	for i, item := range items {
		t.Run(fmt.Sprintf("Test #%d", i+1), func(t *testing.T) {
			pp, err := Parse(item.steamID)
			if err != nil {
				t.Fatalf("%s: %v", item.steamID, err)
			}

			log.Println(pp)
			assert.Equal(t, pp.NickName, item.nickName)
			assert.Equal(t, pp.AvatarUrl[len(pp.AvatarUrl) - 4:], ".jpg")
		})
	}
}
