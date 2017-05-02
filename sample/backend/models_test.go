package backend

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDefaultChannels(t *testing.T) {
	engine := NewEngine()

	output, err := json.MarshalIndent(&ChannelsResponse{
		Channels: engine.AppState.ChannelNameList(),
	}, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Output: " + string(output))

	expected := `{
    "channels": [
        "#general",
        "#random"
    ]
}`
	if expected != string(output) {
		t.FailNow()
	}
}

func TestDefaultUsers(t *testing.T) {
	engine := NewEngine()

	channel := engine.AppState.Channels["#general"]
	output, err := json.MarshalIndent(channel.UserList(engine.AppState), "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Output: " + string(output))

	expected := `[
    {
        "name": "SpeakeasyBot",
        "image_url": "http://localhost:3000/assets/assets/avatars/justice.png"
    }
]`
	if expected != string(output) {
		t.FailNow()
	}
}
