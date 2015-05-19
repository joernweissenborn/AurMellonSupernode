package storage

import (
	"log"
	"github.com/joernweissenborn/aursir4go/aurarath"
	"github.com/joernweissenborn/propertygraph2go/propertygraph"
)
var peernr = 0
type peerproperties struct {
	peer    *aurarath.Peer
	supernode bool
}

type Peer struct {
	g propertygraph.PropertyGraph
	Id string
}

func GetPeer(Id string, g propertygraph.PropertyGraph) Peer {
	return Peer{g,Id}
}

func GetAllPeers(g propertygraph.PropertyGraph, root string) []Peer {
	peers := []Peer{}
	for _, nodeedge := range g.GetVertex(root).Outgoing() {
		if nodeedge.Label() == KNOWN_PEER_EDGE {
			peers = append(peers, GetPeer(nodeedge.Head().Id(),g))
		}
	}
	return peers

}



func (p Peer) Exists() bool {
	return p.g.GetVertex(p.Id)!=nil
}

func (p Peer) IsSuperNode() bool {
	return  p.g.GetVertex(p.Id).Properties().(peerproperties).supernode
}

func (p Peer) GetPeer() (*aurarath.Peer) {
	return  p.g.GetVertex(p.Id).Properties().(peerproperties).peer
}

func (p Peer) Create(peer *aurarath.Peer, supernode bool, root string){
	edge := KNOWN_PEER_EDGE
	if supernode{
		edge = KNOWN_SUPERNODE_EDGE
	}
	peernr++
	log.Println("peerNr",peernr)
	//log.Println("STORAGECORE","Registering Peer ",peer.IdString())
	pv := p.g.CreateVertex(peer.IdString(), peerproperties{peer,supernode})
	p.g.CreateEdge(GenerateUuid(), edge,pv, p.g.GetVertex(root),nil)


	return
}


func (p Peer) Remove(){
	if !p.Exists() {
		return
	}
//	for _, i := range app.GetImports() {
//		i.Remove()
//	}
//	for _, e := range app.GetExports() {
//		e.Remove()
//	}

		log.Println("STORAGECORE","Removing App",p.Id)
		p.g.RemoveVertex(p.Id)
}


/*

func (app App) GetExports() (exports []Export){
	exportids := []string{}
	exports = []Export{}
	c := make(chan string)

	app.agent.Read(func (sc *storage.StorageCore){
		av := sc.GetVertex(app.Id)
		if av!=nil {
			for _, expedge := range av.Outgoing() {
				if expedge.Label() == EXPORT_EDGE {
					c<- expedge.Head().Id()
				}
			}
		}
		close(c)
	})
	for id := range c {
		exportids = append(exportids, id)
	}

	for _, id := range exportids {
		exports = append(exports, GetExportById(id,app.agent))
	}


	return
}

func (app App) GetImports() (imports []Import){
	importids := []string{}
	imports = []Import{}
	c := make(chan string)

	app.agent.Read(func (sc *storage.StorageCore){
		av := sc.GetVertex(app.Id)
		for _, expedge := range av.Outgoing() {
			if expedge.Label() == IMPORT_EDGE {
				c<- expedge.Head().Id()
			}
		}
		close(c)
	})
	for id := range c {
		importids = append(importids, id)
	}
	for _,id := range importids {
		imports = append(imports, GetImportById(id,app.agent))
	}
	return
}
*/
