package backend

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"sort"
	"time"
)

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(t).Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

func (t Timestamp) Unix() int64 {
	return time.Time(t).Unix()
}

func (t Timestamp) UnixNano() int64 {
	return time.Time(t).UnixNano()
}

type User struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type Message struct {
	Id        string    `json:"id"`
	UserName  string    `json:"user_name"`
	CreatedAt Timestamp `json:"created_at"`
	Body      string    `json:"body"`
}

type Channel struct {
	Name     string     `json:"name"`
	Messages []*Message `json:"messages"`
}

type AppState struct {
	Name     string
	Color    string
	Users    map[string]*User
	Channels map[string]*Channel
}

type ChannelsResponse struct {
	Channels []string `json:"channels"`
}

type MessagesResponse struct {
	Users    []*User    `json:"users"`
	Messages []*Message `json:"messages"`
}

type NewMessageRequest struct {
	UserName string `json:"user_name"`
	Body     string `json:"body"`
}

type NewChannelRequest struct {
	UserName    string `json:"user_name"`
	ChannelName string `json:"channel_name"`
}

type ErrorResponse struct {
	Error *Error `json:"error"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func DefaultAppState() *AppState {
	state := &AppState{
		Name:     "Speakeasy",
		Color:    "#FFEC535E",
		Channels: make(map[string]*Channel, 0),
		Users:    make(map[string]*User, 0),
	}
	bot := &User{
		Name: "SpeakeasyBot",
		ImageUrl: GetAsset("assets/avatars/justice.png").Url(),
	}
	state.PutUser(bot)

	general := &Channel{Name: "#general", Messages: []*Message{}}
	state.PutChannel(general)
	random := &Channel{Name: "#random", Messages: []*Message{}}
	state.PutChannel(random)

	general.AddMessage(CreateMessage(bot.Name, "Hi! Welcome to Speakeasy 👋"))

	return state
}

func (state *AppState) PutChannel(channel *Channel) {
	if state.Channels == nil {
		state.Channels = make(map[string]*Channel)
	}
	state.Channels[channel.Name] = channel
}

func (state *AppState) PutUser(user *User) {
	if state.Users == nil {
		state.Users = make(map[string]*User)
	}
	state.Users[user.Name] = user
}

func (channel *Channel) AddMessage(message *Message) {
	channel.Messages = append(channel.Messages, message)
}

func (channel *Channel) LastMessage() *Message {
	return channel.Messages[len(channel.Messages)-1]
}

func (channel *Channel) UserList(state *AppState) []*User {

	userMap := make(map[string]*User)
	for _, message := range channel.Messages {
		userMap[message.UserName] = state.Users[message.UserName]
	}
	userList := make([]*User, 0)
	for _, user := range userMap {
		userList = append(userList, user)
	}

	return userList
}

func (state *AppState) ChannelList() []*Channel {
	channelMap := state.Channels

	channelList := make([]*Channel, 0)
	for _, channel := range channelMap {
		channelList = append(channelList, channel)
	}
	return channelList
}

func (state *AppState) ChannelNameList() []string {
	channelMap := state.Channels

	nameList := make([]string, 0)
	for _, channel := range channelMap {
		nameList = append(nameList, channel.Name)
	}
	sort.Strings(nameList)
	return nameList
}

func CreateMessage(userName string, body string) *Message {
	messageId, _ := uuid.NewV4()
	return &Message{
		Id:        messageId.String(),
		UserName:  userName,
		CreatedAt: Timestamp(time.Now()),
		Body:      body,
	}
}
