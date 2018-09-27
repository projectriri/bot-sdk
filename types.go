package botsdk

type Message interface {
	MessageID() int64
	Type() string
	Text() string
	Chat() Chat
	From() User
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
}

type Update interface {
	UpdateID() int64
	Bot() Bot
	Chat() Chat
	From() User
	Type() string
	Text() string
	IsMessage() bool
	Message() Message

	Reply(config MessageConfig, message interface{}) (interface{}, error)
}

// UpdateChannel is a channel for getting update.
type UpdateChannel <-chan Update

// UpdateErrorChannel is a channel for errors when getting update.
type UpdateErrorChannel <-chan error
