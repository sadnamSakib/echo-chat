package models

// ChatMessage represents a chat message
type Message struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

type Room struct {
	RoomId   string    `json:"roomId"`
	Users    []string  `json:"users"`
	Messages []Message `json:"messages"`
}
