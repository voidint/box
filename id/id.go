// Copyright (c) 2025 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package id

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

// defaultNode is the singleton instance of snowflake ID generator node
var defaultNode *Node

func init() {
	defaultNode = MustNewNode(int64(
		rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1024), // Random integer in [0, 1024) range for node identifier
	))
}

// Int64 generates a snowflake ID as int64 from default node
func Int64() int64 {
	return defaultNode.Int64()
}

// String generates a snowflake ID as base58 string from default node
func String() string {
	return defaultNode.String()
}

// Node represents a snowflake ID generator node
type Node struct {
	id   int64
	core *snowflake.Node
}

// NewNode creates a snowflake ID generator node.
// nodeID must be in the range [0, 1023] as per snowflake specifications
func NewNode(nodeID int64) (*Node, error) {
	core, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}
	return &Node{
		id:   nodeID,
		core: core,
	}, nil
}

// MustNewNode creates a snowflake ID generator node.
// Panics if nodeID is outside the valid range [0, 1023]
func MustNewNode(nodeID int64) *Node {
	n, err := NewNode(nodeID)
	if err != nil {
		panic(err)
	}
	return n
}

// Int64 generates a snowflake ID as int64 from this node
func (n *Node) Int64() int64 {
	return n.core.Generate().Int64()
}

// String generates a snowflake ID as base58 string from this node
func (n *Node) String() string {
	return n.core.Generate().String()
}
