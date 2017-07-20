package models

type Message struct {
	MessageID     string
	MessageOwner  User
	Type          string
	Text          string
	sentTimestamp int64
	readTimestamp int64
}
