package boot

import (
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/aursir4go/aurmellon"
	"github.com/joernweissenborn/stream2go"
)

func BootAurMellon(n *aurarath.Node) stream2go.Stream {
	return n.RegisterProtocol(aurmellon.AurMellon{})
}

