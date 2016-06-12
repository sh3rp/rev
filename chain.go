package rev

type Chain struct {
	RefId []byte
	root  *Node
	count int
}

func NewChain(key []byte, data []byte) *Chain {
	return &Chain{
		RefId: key,
		root: &Node{
			parentHash: []byte{0},
			hash:       hash([]byte{0}, key, data),
			Data:       data,
		},
		count: 1,
	}
}

func (chain *Chain) Revise(value []byte, metadata map[string]string) *Node {
	last := chain.Last()

	if last != nil {
		if eq(last.Data, value) {
			return last
		}
	}

	var parentHash []byte

	if chain.root == nil {
		parentHash = []byte{0}
	} else {
		last := chain.Last()
		if last != nil {
			parentHash = last.hash
		}
	}

	node := &Node{}

	node.parentHash = parentHash
	node.hash = hash(parentHash, chain.RefId, value)
	node.Data = value
	node.Metadata = metadata

	if chain.root == nil {
		chain.root = node
		chain.count = 1
	} else {
		ptr := chain.root
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = node
		node.prev = ptr
		chain.count = chain.count + 1
	}
	return node
}

func (chain *Chain) IsValid() bool {
	ptr := chain.root

	for ptr != nil {
		if eq(ptr.hash, hash(ptr.parentHash, chain.RefId, ptr.Data)) {
			return false
		}
		ptr = ptr.next
	}

	return true
}

func (chain *Chain) Last() *Node {
	ptr := chain.root
	var last *Node
	for ptr != nil {
		last = ptr
		ptr = ptr.next
	}

	return last
}

func (chain *Chain) Nodes() []*Node {
	var nodes []*Node

	ptr := chain.root

	for ptr != nil {
		nodes = append(nodes, ptr)
		ptr = ptr.next
	}

	return nodes
}
