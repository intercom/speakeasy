package backend

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (engine *SampleEngine) channelsResponse(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ChannelsResponse{engine.AppState.ChannelNameList()})
}

func (engine *SampleEngine) newChannelResponse(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newChannel NewChannelRequest
	err := decoder.Decode(&newChannel)
	if err != nil {
		panic("Could not parse request")
	}

	engine.AppState.PutChannel(&Channel{Name: newChannel.ChannelName, Messages: []*Message{}})
}

func (engine *SampleEngine) channelByNameResponse(w http.ResponseWriter, r *http.Request) {
	channelName := mux.Vars(r)["channelName"]
	channel := engine.AppState.Channels[channelName]
	if channel == nil {
		panic("Could not find channel: '" + channelName + "'")
	} else {
		json.NewEncoder(w).Encode(&MessagesResponse{
			Users:    channel.UserList(engine.AppState),
			Messages: channel.Messages,
		})
	}
}

func (engine *SampleEngine) replyResponse(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var reply NewMessageRequest
	err := decoder.Decode(&reply)
	if err != nil {
		panic("Could not parse request")
	}

	channelName := mux.Vars(r)["channelName"]
	if engine.AppState.Users[reply.UserName] == nil {
		engine.AppState.PutUser(&User{Name: reply.UserName})
	}
	message := CreateMessage(reply.UserName, reply.Body)
	engine.AppState.Channels[channelName].AddMessage(message)
	json.NewEncoder(w).Encode(message)
}

func (engine *SampleEngine) assetResponse(w http.ResponseWriter, r *http.Request) {
	assetPath := mux.Vars(r)["assetDirectory"] + "/" + mux.Vars(r)["assetFilename"]
	asset := GetAsset(assetPath)

	if asset != nil {
		data := asset.Data()
		if data == nil {
			panic("Couldn't load image data")
		}

		w.Header().Set("Content-Type", asset.MimeType)
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))

		if _, err := w.Write(data); err != nil {
			panic("unable to write asset.")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (engine *SampleEngine) emptyResponse(w http.ResponseWriter, r *http.Request) {
}
