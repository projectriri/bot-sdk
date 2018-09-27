package botsdk

type UpdateConfig struct {
	BufferSize      int
	PreloadUserInfo bool
}

type SendConfig struct {
	MessageConfig
	ChatConfig
}

const (
	TypeText = iota
	TypeImage
	TypeSticker // TG Only
	TypeRecord  // QQ Only
	TypeDeleteMessage
	TypeEditMessage // TG Only
)

const (
	FormatPlainText   = iota
	FormatMarkdown    // TG Only
	FormatHTML        // TG Only
	FormatRawCQString // QQ Only
)

const (
	UploadModeBase64 = iota
	UploadModeLocal
	UploadModeWeb
)

type MessageConfig struct {
	MediaType         int
	ParseMode         int
	ReplyToMessageID  int64 // TG Only
	DisableWebPreview bool  // TG Only
	UploadMode        int   // QQ Only
	NoCache           bool  // QQ Only (UploadModeWeb)
}

type ChatConfig struct {
	Messenger string
	ChatID    int64
	ChatType  string
	Title     string
}
