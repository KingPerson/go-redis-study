package sortedset

import (
	"math/rand"
	"time"
)

const (
	Zskiplist_MAXLEVEL = 64
	Zskiplist_P        = 0.25
)

//定义skiplist Node 结构

type Element struct {
	Ele   string
	Score float64
}

type ZskiplistNodElevel struct {
	forward *ZskiplistNode //指向本层下一个节点
	span    int64          //forward指向的节点与本节点之间的元素个数
}

type ZskiplistNode struct {
	Element
	backword *ZskiplistNode
	level    []*ZskiplistNodElevel
}

type Zskiplist struct {
	header, tail *ZskiplistNode
	length       int64
	level        int16
}

func ZslCreateNode(level int16, score float64, ele string) (zn *ZskiplistNode) {
	zn = &ZskiplistNode{
		Element: Element{
			Ele:   ele,
			Score: score,
		},
		backword: nil,
		level:    make([]*ZskiplistNodElevel, level),
	}
	for i := range zn.level {
		zn.level[i] = new(ZskiplistNodElevel)
	}
	return zn
}

func ZslCreate() (zsl *Zskiplist) {
	zsl = &Zskiplist{
		header: ZslCreateNode(Zskiplist_MAXLEVEL, 0, ""),
		tail:   nil,
		length: 0,
		level:  1,
	}
	return zsl
}

func ZslRandomLevel() (level int16) {
	level = 1
	rand.Seed(time.Now().UnixNano())
	for float32(rand.Int31()&0xFFFF) < (Zskiplist_P * 0xFFFF) {
		level++
	}
	if level < Zskiplist_MAXLEVEL {
		return level
	}
	return Zskiplist_MAXLEVEL
}

func (zskiplist *Zskiplist) ZslInsert(ele string, score float64) *ZskiplistNode {
	//1.Find the location to insert
	//2.create level and update level update,rank
	//3.create ZskiplistNode
	//4.make node and link into skiplist
	//5.update span covered by update[i] as node is inserted here
	//6.increment span for untouched levels
	//7.set backward node
	//8.set node.level[0].forward.backword
	//9.set zskiplist tail

	update := make([]*ZskiplistNode, Zskiplist_MAXLEVEL)
	rank := make([]int64, Zskiplist_MAXLEVEL)

	x := zskiplist.header
	for i := zskiplist.level - 1; i >= 0; i-- {
		if i == zskiplist.level-1 {
			rank[i] = 0
		} else {
			rank[i] += rank[i+1]
		}
		if x.level[i] != nil {
			for x.level[i].forward != nil &&
				(x.level[i].forward.Score < score ||
					(x.level[i].forward.Score == score && x.level[i].forward.Ele != ele)) {

				rank[i] += x.level[i].span
				x = x.level[i].forward
			}
		}

		update[i] = x
	}

	level := ZslRandomLevel()
	if level > zskiplist.level {
		for i := zskiplist.level; i < level; i++ {
			rank[i] = 0
			update[i] = zskiplist.header
			update[i].level[i].span = zskiplist.length
		}
	}

	node := ZslCreateNode(level, score, ele)
	for i := int16(0); i < level; i++ {
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node

		node.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < zskiplist.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == zskiplist.header {
		node.backword = nil
	} else {
		node.backword = update[0]
	}

	if node.level[0].forward != nil {
		node.level[0].forward.backword = node
	} else {
		zskiplist.tail = node
	}

	zskiplist.length++
	return node
}

func (zskiplist *Zskiplist) ZslDeleteNode(node *ZskiplistNode, update []*ZskiplistNode) {
	for i := int16(0); i < zskiplist.level; i++ {
		if update[i].level[i].forward == node {
			update[i].level[i].span += node.level[i].span - 1
			update[i].level[i].forward = node.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}

	if node.level[0].forward != nil {
		node.level[0].forward.backword = node.backword
	} else {
		zskiplist.tail = node.backword
	}

	for zskiplist.level > 1 && zskiplist.header.level[zskiplist.level-1].forward == nil {
		zskiplist.level--
	}
	zskiplist.length--
}

func (zskiplist *Zskiplist) ZslDelete(ele string, score float64) bool {
	update := make([]*ZskiplistNode, Zskiplist_MAXLEVEL)

	x := zskiplist.header
	for i := zskiplist.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.Score < score ||
				(x.level[i].forward.Score == score && x.level[i].forward.Ele != ele)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	node := x.level[0].forward
	if node != nil && score == node.Score && ele == node.Ele {
		zskiplist.ZslDeleteNode(node, update)
		// free node
		return true
	}
	return false
}
