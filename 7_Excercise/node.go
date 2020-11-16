package node

import (
	"time"
	"math/rand"
)

const(
	electionTimeoutMinMs = 1500                    // min. milliseconds to wait for a heartbeat (starts election when timing out)
	electionTimeoutMaxMs = 3000                    // max. milliseconds to wait for a heartbeat (starts election when timing out)
	heartbeatTicker      = 1000 * time.Millisecond // time between two heartbeats
)

type Node struct{
	id int
	state State
	currentTerm int
	voteReqChannel []chan int
	voteResChannel []chan bool
	heartbeatReceived chan int
	electionTimer *time.Timer // when timeout is reached, a new election begins
	heartbeatTicker *time.Timer // when timeout is reached, a new election begins
	currentLeader int
	votedFor int
	votedBy int
}

func NewNode(id int) *Node {
	node := new(Node)
	node.id = id
	node.currentTerm = 0
	node.votedFor = -1
	node.state = State.FOLLOWER
	node.votedBy = 0
	node.electionTimer = time.NewTimer(randomElectionTimeout())
	node.heartbeatTicker = time.NewTicker(heartbeatTicker)
	node.heartbeatTicker.Stop()
	node.voteReqChannel[0] := make(chan int)
	node.voteReqChannel[1] := make(chan int)
	node.voteReqChannel[2] := make(chan int)
	node.voteReqChannel[3] := make(chan int)
	node.voteReqChannel[4] := make(chan int)

	node.voteResChannel[0] := make(chan bool,1)
	node.voteResChannel[1] := make(chan bool,1)
	node.voteResChannel[2] := make(chan bool,1)
	node.voteResChannel[3] := make(chan bool,1)
	node.voteResChannel[4] := make(chan bool,1)
	return node

	
}

func randomElectionTimeout() time.Duration {
	ms := rand.Intn(electionTimeoutMaxMs-electionTimeoutMinMs) + electionTimeoutMinMs
	return time.Duration(ms) * time.Millisecond
}


func (n *Node) StartVoting()
{
	n.electionTimer.Reset(randomElectionTimeout())
	n.State = State.CANDIDATE
	n.votedFor = n.Id
	n.votedBy = 1
	for _, req := range n.voteReqChannel {
		req <- n.id
	}
	for _, res := range n.voteResChannel {
		if <- res{
			n.votedBy++
		}
	}
	if n.votedBy > 2{
		n.State = State.LEADER
		n.electionTimer.Stop()
		n.heartbeatTicker.Reset(heartbeatTicker)
	}
}

func FanIn(channels []chan int) chan int {
	output := make(chan int)
	for i := 0; i < len(channels); i++ {
		// fan in
		go func(i int) {
			for {
				n, ok := <-channels[i]
				if !ok {
					break
				}
				output <- n
			}
		}(i)
	}
	return output
}

func (n *Node) SendHeartbeats(){

}

func (n *Node) Run(){
	fanInedReqChannel := FanIn(n.voteReqChannel)
	for{
		select{
		case voteReq <- fanInedReqChannel:
			if n.votedFor == -1{
				n.votedFor = voteReq
				n.voteResChannel[voteReq] <- true
			}
			else{
				n.voteResChannel[voteReq] <- false
			}
		case currentLeader <- n.heartbeatReceived:
			n.currentLeader = currentLeader
			n.heartbeatTicker.Stop()
			n.State = State.FOLLOWER
			n.electionTimer.Reset(randomElectionTimeout())
		case <- n.electionTimer.C:
			n.StartVoting()
		case <- n.heartbeatTicker.C:
			n.SendHeartbeats()
		}
	}
}