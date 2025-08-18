package go_effect_sort

func BinaryInsert[T any](oldSlice []T, val T, equalOrder EqualOrderExternal[T], greaterThan GreaterThanExternal[T], move Move[T]) (newSlice []T, insertIdx int) {
	insertIdx = BinaryInsertIndex(oldSlice, val, equalOrder, greaterThan)
	newSlice = append(oldSlice, val)
	i := len(newSlice) - 2
	for i >= insertIdx {
		move(newSlice, i, i+1)
		i -= 1
	}
	newSlice[insertIdx] = val
	return
}

func BinaryInsertIndex[T any](oldSlice []T, val T, equalOrder EqualOrderExternal[T], greaterThan GreaterThanExternal[T]) (insertIdx int) {
	var lo int = 0
	var hi int = len(oldSlice)
	insertIdx = lo
	for {
		if hi <= lo {
			goto setIdxToLo
		}
		insertIdx = lo + ((hi - lo) >> 1)
		if equalOrder(oldSlice, insertIdx, val) {
			goto returnIdx
		}
		if greaterThan(oldSlice, insertIdx, val) {
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

func BinarySearch[T any](slice []T, val T, equalValue EqualValueExternal[T], greaterThan GreaterThanExternal[T]) (idx int, found bool) {
	var lo int = 0
	var hi int = len(slice)
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

func BinaryAlter[T any](slice []T, alteredIdx int, equalOrder EqualOrderExternal[T], greaterThan GreaterThanExternal[T], move Move[T]) (newIdx int) {
	newIdx = BinaryInsertIndex(slice, slice[alteredIdx], equalOrder, greaterThan)
	if newIdx == alteredIdx {
		return
	}
	val := slice[alteredIdx]
	if newIdx < alteredIdx {
		i := alteredIdx - 1
		for i >= newIdx {
			move(slice, i, i+1)
			i -= 1
		}
	} else {
		i := alteredIdx + 1
		for i <= newIdx {
			move(slice, i, i-1)
			i += 1
		}
	}
	slice[newIdx] = val
	return
}
