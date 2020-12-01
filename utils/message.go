package utils

import (
	"encoding/json"
)

func FormatMessage(content interface{}) ([]byte, error) {
	switch content := content.(type) {
	case string:
		content = map[string]string{"content": content}

		// TODO : Add support for attachments
	}

	b, err := json.Marshal(content)

	return b, err
}
