package boot

import (
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/propertygraph2go/transcator"
	"github.com/joernweissenborn/stream2go"
	"github.com/joernweissenborn/aursir4go/aurmellon"
	"AurMellonSupernode/processors"
)

func BootProcessors(n *aurarath.Node, in stream2go.Stream, t transcator.Transcator) {
	t.Transaction(createRoot(n))
	in.Where(aurmellon.IsHelloIamNodeMessage).Listen(processors.AddPeer(t,n))
	in.Where(aurmellon.IsHelloIamSuperNodeMessage).Listen(processors.AddSupernode(t,n))
}

func createRoot(n *aurarath.Node) transcator.TransactionFunc {
	return func(t *transcator.Transaction) {
		t.CreateVertex(n.Self.IdString(), n.Self)
	}
}
