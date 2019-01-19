package network

import (
	"errors"
	"log"
	"net"
	"time"
)

// ConnMgr 保存map映射
type ConnMgr struct {
	Listener    net.Listener
	LocalIP     string
	Peers       map[string]*Peer
	ServerPeers map[string]*ServerPeer
}

func (connMgr *ConnMgr) Start() {
	err := connMgr.listen(connMgr.LocalIP)
	if err != nil {
		log.Fatalln("Error: Faild to start Tcp listen")
	}
	go connMgr.acceptConn()
}

// StartListen 开始监听某个端口
func (connMgr *ConnMgr) listen(localIP string) error {
	connMgr.LocalIP = localIP
	connMgr.Peers = make(map[string]*Peer)
	listen, err := net.Listen("tcp", connMgr.LocalIP)
	if err != nil {
		log.Printf("开始监听时出现错误错误：%v", err)
		return errors.New("Faild to start listner")
	}
	connMgr.Listener = listen
	return nil
}

// RequestConn 开始发起一个连接
func (connMgr *ConnMgr) RequestConn(targetIP string) *Peer {
	var conn net.Conn
	var err error
	for {
		conn, err = net.DialTimeout("tcp", targetIP, time.Second)
		if err != nil {
			//log.Printf("建立连接不成功：%v", err)
		} else {
			break
		}
	}
	log.Printf("成功建立连接：%s", conn.RemoteAddr().String())
	ip := conn.RemoteAddr().String()
	peer := &Peer{ip, conn}
	log.Printf("创建一个peer, peer ip:%s", peer.targetIP)
	connMgr.Peers[ip] = peer
	return peer

}

// AcceptConn 接收一个连接
func (connMgr *ConnMgr) acceptConn() {
	for {
		log.Println("正在Accept")
		conn, err := connMgr.Listener.Accept()
		if err != nil {
			log.Printf("接收连接时发生错误：%v", err)
			return
		}
		log.Printf("成功建立连接：%s", conn.RemoteAddr().String())
		ip := conn.RemoteAddr().String()
		peer := &Peer{ip, conn}
		log.Printf("创建一个peer, peer ip:%s", peer.targetIP)
		connMgr.Peers[ip] = peer
		go peer.RecvMsg()
	}
}

func (connMgr *ConnMgr) BroadCast(msg Message) {
	for _, v := range connMgr.Peers {
		v.SendMsg(msg)
	}
}
