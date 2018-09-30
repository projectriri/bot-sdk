package botsdk

type SendConfig struct {
	MessageConfig
	ChatConfig
}

const (
	TypeText = iota
	TypeImage
	TypeSticker
	TypeRecord
	TypeDeleteMessage
	TypeEditMessage
)

const (
	FormatPlainText = iota
	FormatMarkdown
	FormatHTML
	FormatRawCQString
)

type MessageConfig struct {
	MediaType         int
	ParseMode         int
	MessageID         int64
	ReplyToMessageID  int64
	DisableWebPreview bool
	NoCache           bool
}

type ChatConfig struct {
	ChatID   int64
	ChatType string
	Title    string
}
