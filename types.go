package botsdk

type Message interface {
	MessageID() int64
	Bot() Bot
	Type() string
	Text() string
	Chat() Chat
	From() User
	Reply(config MessageConfig, message interface{}) (interface{}, error)
}

type Chat interface {
	ChatConfig() ChatConfig
}

type User interface {
	Messenger() string
	UserID() int64
	UserName() string
	DisplayName() string
	FirstName() string
	LastName() string
	PrivateChat() ChatConfig
	At(parseMode int) string
}

type Update interface {
	UpdateID() int64
	Bot() Bot
	IsMessage() bool
	Message() Message
}

// UpdateChannel is a channel for getting update.
type UpdateChannel <-chan Update

// UpdateErrorChannel is a channel for errors when getting update.
type UpdateErrorChannel <-chan error
