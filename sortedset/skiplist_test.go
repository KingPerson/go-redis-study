package sortedset

import (
	"fmt"
	"testing"
)

func TestZskiplist_ZslInsert(t *testing.T) {
	zsl := ZslCreate()
	node1 := zsl.ZslInsert("song", 100)
	fmt.Println(node1)

	node2 := zsl.ZslInsert("jia", 200)
	fmt.Println(node2)

	fmt.Println(zsl)

	zsl.ZslDelete(node1.Ele, node1.Score)

	fmt.Println(node1)

	zsl.ZslDelete(node2.Ele, node2.Score)

	fmt.Println(zsl)
}


func TestZskiplist_ZslDelete(t *testing.T) {

}