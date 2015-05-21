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
		m,_ := aurarath.ToIncomingProtocolMessage(d)
		p := m.Sender

		t.Transaction(addpeer(p,false,n.Self.Id.String()))
	}
}
func AddSupernode(t transcator.Transcator, n *aurarath.Node) stream2go.Suscriber {
	return func(d interface {}){
		m,_ := aurarath.ToIncomingProtocolMessage(d)
		p := m.Sender
		t.Transaction(addpeer(p,true,n.Self.Id.String()))
		p.Send(aurmellon.HelloIamSuperNodeMessage{})
	}
}

func addpeer(p *aurarath.Peer,supernode bool, root string) transcator.TransactionFunc {
	return func(t *transcator.Transaction){
		if t==nil {
			panic("transaction nil")
		}
		peer := storage.GetPeer(p.Id.String(),t)
		if !peer.Exists() {
			peer.Create(p,supernode,root)
		}
	}
}
