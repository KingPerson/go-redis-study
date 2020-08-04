package sortedset

type Sortedset struct {
	dict      map[string]*Element
	zskiplist *Zskiplist
}

func create() *Sortedset {
	return &Sortedset{
		dict:      make(map[string]*Element),
		zskiplist: ZslCreate(),
	}
}

func (sortedset *Sortedset) Add(ele string, score float64) bool {
	return true
}

func (sortedset *Sortedset) Delete(ele string) bool {
	return true
}

//ZSCORE
func (sortedset *Sortedset) Get(ele string) (element *Element, ok bool) {
	return element, true
}

//ZCARD
func (sortedset *Sortedset) Len() int64 {
	return int64(0)
}

//ZRANK && ZREVRANK
func (sortedset *Sortedset) GetRank(ele string, desc bool) (rank int64) {
	return rank
}
