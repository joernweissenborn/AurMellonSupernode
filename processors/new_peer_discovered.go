package processors

import (
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/stream2go"
	"github.com/joernweissenborn/aursir4go/aurmellon"
	"log"
	"time"
)
var discpeers = 0
func NewPeerDiscovered(n *aurarath.Node) stream2go.Suscriber{
	return func(d interface {}) {
		p := d.(*aurarath.Peer)
		time.Sleep(1 * time.Millisecond)
		discpeers++
		log.Println("DIS PEERS",discpeers)
		sendToPeer(n,p,aurmellon.HelloIamSuperNodeMessage{})
	}
}

