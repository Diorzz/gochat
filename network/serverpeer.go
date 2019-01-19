package network

// ServerPeer 应用层逻辑
// peer 当前peer连接
import (
	"log"
)

type ServerPeer struct {
	peer *Peer
}

func NewServerPeer(peer *Peer) *ServerPeer {
	sp := &ServerPeer{
		peer: peer,
	}
	return sp
}

func (sp *ServerPeer) Send(msg Message) {
	msgDecode := MsgToRlpData(msg)
	sp.peer.Send(msgDecode)
}

func (sp *ServerPeer) Read() {

	for {
		data := sp.peer.Read()
		msg, err := DataToRlpMsg(data)
		if err != nil {
			log.Println(err)
		}
		log.Println(msg)
	}
}
