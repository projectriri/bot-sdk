package botsdk

type Bot interface {
	GetUpdatesChan(config UpdateConfig) (UpdateChannel, UpdateErrorChannel)

	Send(config SendConfig, message interface{}) (messageID int64, err error)
	Delete(config ChatConfig, messageID int64) error

	At(User) string
	Self() User
}
