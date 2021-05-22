package pass

import (
	"github.com/MihaiBlebea/go-pass-client/caller"
	"github.com/MihaiBlebea/go-pass-client/resource/catalog"
)

func New(token string) catalog.Service {
	return catalog.New(caller.New(token))
}
