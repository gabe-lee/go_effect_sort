package go_effect_sort

import (
	ll "github.com/gabe-lee/go_list_like"
)

type IndexMap[T any, I Index] struct {
	GreaterThan func(slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool
	LessThan    func(slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool
	Indexes     ll.ListLike[I]
}

func NewIndexMaps[T any, I Index](maps ...IndexMap[T, I]) []IndexMap[T, I] {
	return maps
}

func NewIndexMap[T any, I Index](indexes ll.ListLike[I], greaterThan func(slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool, lessThan func(slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool) IndexMap[T, I] {
	return IndexMap[T, I]{
		GreaterThan: greaterThan,
		LessThan:    lessThan,
		Indexes:     indexes,
	}
}

func (m IndexMap[T, I]) ShiftIndexUp(list ll.ListLike[T], oldMovedIndex, newMovedIndex I) (mappedIdx int, found bool) {
	var maxCount I = newMovedIndex - oldMovedIndex
	var count I = 0
	found = false
	for i := range m.Indexes.Len() {
		v := m.Indexes.Get(i)
		if v == oldMovedIndex {
			found = true
			mappedIdx = i
			if count >= maxCount && found {
				break
			}
		} else if v > oldMovedIndex && v <= newMovedIndex {
			ll.SetSubtract(m.Indexes, i, 1)
			count += 1
			if count >= maxCount && found {
				break
			}
		}
	}
	if found {
		m.Indexes.Set(mappedIdx, newMovedIndex)
	}
	return
}
func (m IndexMap[T, I]) MoveIndexesUp(list ll.ListLike[T], startIdx I) {
	var maxCount I = I(list.Len()) - startIdx
	var count I = 0
	for i := range m.Indexes.Len() {
		v := m.Indexes.Get(i)
		if v >= startIdx {
			ll.SetAdd(m.Indexes, i, 1)
			count += 1
			if count >= maxCount {
				break
			}
		}
	}
}

func (m IndexMap[T, I]) MoveIndexesDown(list ll.ListLike[T], startIdx I) {
	var maxCount I = I(list.Len()) - startIdx
	var count I = 0
	for i := range m.Indexes.Len() {
		v := m.Indexes.Get(i)
		if v >= startIdx {
			ll.SetSubtract(m.Indexes, i, 1)
			count += 1
			if count >= maxCount {
				break
			}
		}
	}
}

func (m IndexMap[T, I]) ShiftIndexDown(list ll.ListLike[T], oldMovedIndex, newMovedIndex I) (mappedIdx int, found bool) {
	var maxCount I = oldMovedIndex - newMovedIndex
	var count I = 0
	found = false
	for i := range m.Indexes.Len() {
		v := m.Indexes.Get(i)
		if v == oldMovedIndex {
			found = true
			mappedIdx = i
			if count >= maxCount && found {
				break
			}
		} else if v < oldMovedIndex && v >= newMovedIndex {
			ll.SetAdd(m.Indexes, i, 1)
			count += 1
			if count >= maxCount && found {
				break
			}
		}
	}
	if found {
		m.Indexes.Set(mappedIdx, newMovedIndex)
	}
	return
}

func (m IndexMap[T, I]) ResortAltered(list ll.ListLike[T], mapIdx int) {
	val := m.Indexes.Get(mapIdx)
	for mapIdx > 0 && m.LessThan(list, m.Indexes, mapIdx, mapIdx-1) {
		ll.Move(m.Indexes, mapIdx-1, mapIdx)
		mapIdx -= 1
	}
	for mapIdx < m.Indexes.Len()-1 && m.GreaterThan(list, m.Indexes, mapIdx, mapIdx+1) {
		ll.Move(m.Indexes, mapIdx+1, mapIdx)
		mapIdx += 1
	}
	m.Indexes.Set(mapIdx, val)
}
