package model

type Room struct {
	ID        int
	Messages  []Message
	Observers map[int]struct {
		UserID  int
		Message chan *Message
	}
}
