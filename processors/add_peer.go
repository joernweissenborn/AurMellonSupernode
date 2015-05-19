package processors

import (
	"github.com/joernweissenborn/stream2go"
	"github.com/joernweissenborn/propertygraph2go/transcator"
	"github.com/joernweissenborn/aursir4go/aurarath"
	"AurMellonSupernode/storage"
	"github.com/joernweissenborn/aursir4go/aurmellon"
)

func AddPeer(t transcator.Transcator, n *aurarath.Node) stream2go.Suscriber {
	return func(d interface {}){
		m,_ := aurarath.ToMessage(d)
		p := n.GetPeer(m.Sender)
		t.Transaction(addpeer(p,true,n.Self.IdString()))
	}
}
func AddSupernode(t transcator.Transcator, n *aurarath.Node) stream2go.Suscriber {
	return func(d interface {}){
		m,_ := aurarath.ToMessage(d)
		p := n.GetPeer(m.Sender)
		t.Transaction(addpeer(p,true,n.Self.IdString()))
		sendToPeer(n,p,aurmellon.HelloIamSuperNodeMessage{})
	}
}

func addpeer(p *aurarath.Peer,supernode bool, root string) transcator.TransactionFunc {
	return func(t *transcator.Transaction){
		if t==nil {
			panic("transaction nil")
		}
		peer := storage.GetPeer(p.IdString(),t)
		if !peer.Exists() {
			peer.Create(p,supernode,root)
		}
	}
}
