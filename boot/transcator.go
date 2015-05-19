package boot

import (
	"github.com/joernweissenborn/propertygraph2go/transcator"
	"github.com/joernweissenborn/propertygraph2go/inmemorygraph"
)

func BootTransactor() transcator.Transcator {
	return transcator.New(inmemorygraph.New())
}
