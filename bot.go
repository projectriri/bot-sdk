package botsdk

type Bot interface {
	GetUpdatesChan(bufferSize int) (UpdateChannel, UpdateErrorChannel, error)

	Send(config SendConfig, message interface{}) (messageID int64, err error)

	Self() (User, error)
}
