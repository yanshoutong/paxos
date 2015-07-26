package paxos

import "log"

func CreateNetwork(nodes ...int) *network {
	nt := network{recvQueue: make(map[int]chan message, 0)}

	for _, node := range nodes {
		nt.recvQueue[node] = make(chan message, 1024)
	}

	return &nt
}

type network struct {
	recvQueue map[int]chan message
}

func (n *network) sendTo(from int, m message) {
	log.Println("Send msg from:", m.from, " send to", m.to, " val:", m.val)
	n.recvQueue[m.to] <- m
}

func (n *network) recevFrom(id int) message {
	retMsg := <-n.recvQueue[id]
	log.Println("Recev msg from:", retMsg.from, " send to", retMsg.to, " val:", retMsg.val)
	return retMsg
}

type nodeNetwork struct {
	id  int
	net *network
}

func (n *nodeNetwork) send(m message) {
	n.net.sendTo(m.to, m)
}

func (n *nodeNetwork) recev() message {
	return n.net.recevFrom(n.id)
}