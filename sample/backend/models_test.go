package backend

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDefaultChannels(t *testing.T) {
	server := NewServer()

	output, err := json.MarshalIndent(&ChannelsResponse{
		Channels: server.AppState.ChannelNameList(),
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
	server := NewServer()

	channel := server.AppState.Channels["#general"]
	output, err := json.MarshalIndent(channel.UserList(server.AppState), "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Output: " + string(output))

	expected := `[
    {
        "name": "SpeakeasyBot",
        "image_url": ""
    }
]`
	if expected != string(output) {
		t.FailNow()
	}
}
