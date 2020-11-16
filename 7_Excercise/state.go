package node

type State int

const (
	FOLLOWER
	LEADER
	CANDIDATE
)

func (s State) String() string {
	switch s {
	case FOLLOWER:
		return "FOLLOWER"
	case LEADER:
		return "LEADER"
	case CANDIDATE:
		return "CANDIDATE"
	default:
		return "Unknown" 
	}
}