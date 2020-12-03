package message

type Attachment struct {
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	URL      string `json:"url"`
	Data     []byte `json:"-"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

func NewAttachment(filename string, url string, data []byte) *Attachment {
	return &Attachment{
		Filename: filename,
		URL:      url,
		Data:     data,
	}
}
