package json2go

import (
	"sort"
)

func getSortedKeys(m map[string]interface{}) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	sort.Sort(ByIDFirst(keys))
	return
}

//ByIDFirst ...
type ByIDFirst []string

func (p ByIDFirst) Len() int { return len(p) }
func (p ByIDFirst) Less(i, j int) bool {
	if p[i] == "id" {
		return true
	} else if p[j] == "id" {
		return false
	}
	return p[i] < p[j]
}
func (p ByIDFirst) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
