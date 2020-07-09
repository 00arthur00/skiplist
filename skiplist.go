package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type SkipList struct {
	//header only reserve forwards, the key/value always empty
	Header     *Node
	level, max int
}
type Node struct {
	Key, Value string
	Forward    []*Node
}

func (s *SkipList) randLevel() int {
	const p = 0.25 //1/4
	h := 0
	for h < s.max && rand.Float32() < p {
		h++
	}
	return h
}

func NewSkipList(maxLevel int) *SkipList {
	return &SkipList{Header: &Node{Forward: make([]*Node, maxLevel, maxLevel)}, level: 0, max: maxLevel}
}

func (s *SkipList) getLTPath(key string) []*Node {
	path := make([]*Node, s.level+1, s.max)
	x := s.Header
	for i := s.level; i >= 0; i-- {
		for x.Forward[i] != nil && x.Forward[i].Key < key {
			x = x.Forward[i]
		}
		path[i] = x
	}
	return path
}
func (s *SkipList) Insert(key, value string) {
	path := s.getLTPath(key)

	x := path[len(path)-1]
	if x.Forward[0] != nil && x.Forward[0].Key == key {
		x.Forward[0].Value = value
		return
	}

	height := s.randLevel()
	if height > s.level {
		for i := s.level; i < height; i++ {
			path = append(path, s.Header)
		}
		s.level = height
	}

	node := &Node{Key: key, Value: value, Forward: make([]*Node, height+1, height+1)}
	for i := height; i >= 0; i-- {
		node.Forward[i] = path[i].Forward[i]
		path[i].Forward[i] = node
	}
	return
}

func (s *SkipList) Search(key string) (string, bool) {

	path := s.getLTPath(key)

	x := path[0]

	if x.Forward[0] != nil && x.Forward[0].Key == key {
		return x.Forward[0].Value, true
	}
	return "", false
}

func (s *SkipList) Delete(key string) {
	path := s.getLTPath(key)
	x := path[len(path)-1]
	if x.Forward[0] != nil && x.Forward[0].Key != key {
		return
	}
	for i := s.level; i >= 0; i-- {
		path[i].Forward[i] = x.Forward[i]
	}
	return
}

//internal usage for debug
func (s *SkipList) print() {
	x := s.Header
	// fmt.Println(s.level, s.max)
	count := 0
	fmt.Println("Height:", s.level+1)
	for x.Forward[0] != nil {
		fmt.Println(count, len(x.Forward), x.Forward[0].Key, x.Forward[0].Value)
		x = x.Forward[0]
		count++
	}
}

//internal usage for debug
func (s *SkipList) printPath(path []*Node) {
	fmt.Println("====path=====")
	for _, node := range path {
		fmt.Println(node.Key, node.Value)
	}

	fmt.Println("********************")
}
