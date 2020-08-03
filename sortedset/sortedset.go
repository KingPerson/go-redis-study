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
