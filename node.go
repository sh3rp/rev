package rev

type Node struct {
	prev       *Node
	next       *Node
	parentHash []byte
	hash       []byte
	Data       []byte
	Metadata   map[string]string
}
