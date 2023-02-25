package id

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

var defaultNode *Node // 默认节点

func init() {
	defaultNode = MustNewNode(int64(
		rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1024), // [0,1024)区间范围内的整数
	))
}

// Int64 返回默认节点生成的int64类型的ID
func Int64() int64 {
	return defaultNode.Int64()
}

// String 返回默认节点生成的string类型的ID
func String() string {
	return defaultNode.String()
}

// Node ID生成器节点
type Node struct {
	id   int64
	core *snowflake.Node
}

// NewNode 返回ID生成器节点。注意：nodeID的取值范围为[0,1023]。
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

// MustNewNode 返回ID生成器节点。若入参ID不满足要求[0,1023]，将发生panic。
func MustNewNode(nodeID int64) *Node {
	n, err := NewNode(nodeID)
	if err != nil {
		panic(err)
	}
	return n
}

// Int64 返回节点生成的int64类型的ID
func (n *Node) Int64() int64 {
	return n.core.Generate().Int64()
}

// String 返回节点生成的string类型的ID
func (n *Node) String() string {
	return n.core.Generate().String()
}
