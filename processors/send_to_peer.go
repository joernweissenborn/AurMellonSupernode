package processors

import (
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/stream2go"
	"log"
	"github.com/joernweissenborn/future2go"
)

var try = 0
var snd = 0
func  sendToPeer(n *aurarath.Node,p *aurarath.Peer, m aurarath.ProtocolMessage) {
	try++
	log.Println("trypeer",try)

	f := n.OpenConnection(p)
	f.Then(send(m))
	f.Err(func(err error) (interface{}, error){
		log.Println("SndpeerERR", err)

		return nil, nil
	}).Then(send(m))
}

func send(m aurarath.ProtocolMessage) future2go.CompletionFunc{
	return func(d interface{}) interface{} {
		snd++
		log.Println("Sndpeer",snd)

		d.(stream2go.StreamController).Add(m)
		return nil
	}
}
