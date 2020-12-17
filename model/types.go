package model

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"time"
)

type User struct {
	AvatarUrl   string `json:"avatarUrl"`
	DisplayName string `json:"displayName"`
	UserName    string `json:"userName"`
}
type Message struct {
	From User      `json:"fromUser"`
	Id   string    `json:"id"`
	Sent time.Time `json:"sent"`
	Text string    `json:"text"'`
	Html string    `json:"html"`
}
type Room struct {
	Messages []*Message
	Name     string
}

func RoomFromArchive(archive string) (*Room, error) {
	data, err := ioutil.ReadFile(archive)
	if err != nil {
		return nil, err
	}
	var msgs []*Message
	err = json.Unmarshal(data, &msgs)
	if err != nil {
		return nil, err
	}
	return &Room{
		Messages: msgs,
		Name:     path.Base(archive),
	}, nil
}
