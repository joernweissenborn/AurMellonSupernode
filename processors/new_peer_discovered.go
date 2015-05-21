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
		pid := d.(aurarath.PeerId)
		p := n.GetPeer(pid.String())
		time.Sleep(100 * time.Millisecond)
		discpeers++
		log.Println("DIS PEERS",discpeers,p.Id,p.Addresses())
		if len(p.Addresses())== 0 {
			panic(p.Id)
		}
		p.OpenConnection()
		p.Send(aurmellon.HelloIamSuperNodeMessage{})
	}
}

