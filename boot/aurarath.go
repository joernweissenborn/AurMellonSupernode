package boot

import (
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/aursir4go/aurarath/implementation/zaurarath"
	"AurMellonSupernode/processors"
)

func BootAurArath() *aurarath.Node {

	node := aurarath.NewNode()
	node.RegisterImplementation(new(zaurarath.Implementation))
	node.NewPeers().Listen(processors.NewPeerDiscovered(node))
	return node
}
