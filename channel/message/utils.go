package message

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Seyz123/yalis/channel/message/embed"
	"net/http"
)

func FormatImage(data []byte) string {
	var encoded []byte

	mediaType := http.DetectContentType(data)
	base64.StdEncoding.Encode(encoded, data)

	return fmt.Sprintf("data:%s;base64,%s", mediaType, string(encoded))
}

func FormatMessage(content interface{}) ([]byte, error) {
	switch content.(type) {
	case string:
		content = map[string]string{"content": content.(string)}

	case *embed.Embed:
		content = &embed.MessageEmbed{Embed: content.(*embed.Embed)}
	}

	b, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	return b, nil
}