package go_effect_sort

import (
	ll "github.com/gabe-lee/go_list_like"
)

func Sorted_Insert[T any](list ll.ListLike[T], val T, equalOrder func(list ll.SliceLike[T], idx int, val T) bool, greaterThan func(list ll.SliceLike[T], idx int, val T) bool, move func(list ll.SliceLike[T], oldIdx int, newIdx int)) (insertIdx int) {
	insertIdx = Sorted_InsertIndex(list, val, equalOrder, greaterThan)
	ll.Append(list, val)
	i := list.Len() - 2
	for i >= insertIdx {
		move(list, i, i+1)
		i -= 1
	}
	ll.Set(list, insertIdx, val)
	return
}

func Sorted_InsertWithMaps[T any, I Index](list ll.ListLike[T], val T, equalOrder func(list ll.SliceLike[T], idx int, val T) bool, greaterThan func(list ll.SliceLike[T], idx int, val T) bool, move func(list ll.SliceLike[T], oldIdx int, newIdx int), indexMaps []IndexMap[T, I]) (insertIdx int) {
	insertIdx = Sorted_Insert(list, val, equalOrder, greaterThan, move)
	for _, m := range indexMaps {
		m.MoveIndexesUp(list, I(insertIdx))
		mapAppendIdx := m.Indexes.Len()
		ll.Append(m.Indexes, I(insertIdx))
		m.ResortAltered(list, mapAppendIdx)
	}
	return
}

func Sorted_InsertIndex[T any](list ll.SliceLike[T], val T, equalOrder func(list ll.SliceLike[T], idx int, val T) bool, greaterThan func(list ll.SliceLike[T], idx int, val T) bool) (insertIdx int) {
	var lo int = 0
	var hi int = list.Len()
	insertIdx = lo
	for {
		if hi <= lo {
			goto setIdxToLo
		}
		insertIdx = lo + ((hi - lo) >> 1)
		if equalOrder(list, insertIdx, val) {
			goto returnIdx
		}
		if greaterThan(list, insertIdx, val) {
			hi = insertIdx
		} else {
			lo = insertIdx + 1
		}
	}
setIdxToLo:
	insertIdx = lo
returnIdx:
	return
}

func Sorted_Search[T any](slice ll.SliceLike[T], val T, equalValue func(list ll.SliceLike[T], idx int, val T) bool, greaterThan func(list ll.SliceLike[T], idx int, val T) bool) (idx int, found bool) {
	var lo int = 0
	var hi int = slice.Len()
	var mid int = lo
	for {
		if hi <= lo {
			return 0, false
		}
		mid = lo + ((hi - lo) >> 1)
		if equalValue(slice, mid, val) {
			return mid, true
		}
		if greaterThan(slice, mid, val) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
}

func Sorted_FindMappedIdx[T any, I Index](slice ll.SliceLike[T], valIdx int, indexMap IndexMap[T, I]) (idx int, found bool) {
	var lo int = 0
	var hi int = slice.Len()
	var mid int = lo
	for {
		if hi <= lo {
			return 0, false
		}
		mid = lo + ((hi - lo) >> 1)
		if indexMap.GreaterThan(slice, indexMap.Indexes, mid, valIdx) {
			hi = mid
		} else if indexMap.LessThan(slice, indexMap.Indexes, mid, valIdx) {
			lo = mid + 1
		} else {
			return mid, true
		}
	}
}

func Sorted_Alter[T any](slice ll.SliceLike[T], alterIdx int, equalOrder func(list ll.SliceLike[T], idx int, val T) bool, greaterThan func(list ll.SliceLike[T], idx int, val T) bool, move func(list ll.SliceLike[T], oldIdx int, newIdx int)) (newIdx int) {
	val := ll.Get(slice, alterIdx)
	newIdx = Sorted_InsertIndex(slice, val, equalOrder, greaterThan)
	if newIdx == alterIdx {
		return
	}
	if newIdx < alterIdx {
		i := alterIdx - 1
		for i >= newIdx {
			move(slice, i, i+1)
			i -= 1
		}
	} else {
		i := alterIdx + 1
		for i <= newIdx {
			move(slice, i, i-1)
			i += 1
		}
	}
	ll.Set(slice, newIdx, val)
	return
}

func Sorted_AlterWithMaps[T any, I Index](slice ll.ListLike[T], alterIdx int, equalOrder func(list ll.SliceLike[T], idx int, val T) bool, greaterThan func(list ll.SliceLike[T], idx int, val T) bool, move func(list ll.SliceLike[T], oldIdx int, newIdx int), maps []IndexMap[T, I]) (newIdx int) {
	newIdx = Sorted_Alter(slice, alterIdx, equalOrder, greaterThan, move)
	if newIdx > alterIdx {
		for _, m := range maps {
			mIdx, found := m.ShiftIndexUp(slice, I(alterIdx), I(newIdx))
			if found {
				m.ResortAltered(slice, mIdx)
			}
		}
	} else if newIdx < alterIdx {
		for _, m := range maps {
			mIdx, found := m.ShiftIndexDown(slice, I(alterIdx), I(newIdx))
			if found {
				m.ResortAltered(slice, mIdx)
			}
		}
	} else {
		for _, m := range maps {
			mIdx, found := Sorted_FindMappedIdx(slice, alterIdx, m)
			if found {
				m.ResortAltered(slice, mIdx)
			}
		}
	}
	return
}
