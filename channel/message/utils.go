package message

import (
	"encoding/json"
	"github.com/Seyz123/yalis/channel/message/embed"
)

func FormatMessage(content interface{}) ([]byte, error) {
	switch content.(type) {
	case string:
		content = map[string]string{"content": content.(string)}

	case *embed.Embed:
		content = &embed.MessageEmbed{Embed: content.(*embed.Embed)}

		// TODO : Add support for attachments
	}

	b, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	return b, nil
}