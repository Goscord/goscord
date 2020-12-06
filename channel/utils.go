package channel

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func FormatImage(data []byte) string {
	var encoded []byte

	mediaType := http.DetectContentType(data)
	base64.StdEncoding.Encode(encoded, data)

	return fmt.Sprintf("data:%s;base64,%s", mediaType, string(encoded))
}