package rev

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	chain := NewChain([]byte("key"), []byte("value1"))
	chain.Revise([]byte("value1"), nil)
	chain.Revise([]byte("value2"), nil)
	chain.Revise([]byte("value3"), nil)
	chain.Revise([]byte("value4"), nil)

	nodes := chain.Nodes()

	for _, n := range nodes {
		fmt.Printf("%x: %v\n", n.hash, string(n.Data))
	}

	assert.True(t, chain.IsValid())
}
