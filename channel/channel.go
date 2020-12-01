package channel

import "github.com/Seyz123/yalis/rest"

type Channel struct {
	Rest *rest.RestClient `json:"-"`
}