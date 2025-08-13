package go_effect_sort

func BinaryInsert[T any](oldSlice []T, val T, equalOrder EqualOrderExternal[T], greaterThan GreaterThanExternal[T], move Move[T]) (newSlice []T, insertIdx int) {
	newSlice = append(oldSlice, val)
	var lo int = 0
	var hi int = len(newSlice) - 1
	var mid int = lo
	for {
		if hi <= lo {
			goto setMidToLo
		}
		mid = lo + ((hi - lo) >> 1)
		if equalOrder(oldSlice, mid, val) {
			goto moveUpAndIns
		}
		if greaterThan(newSlice, mid, val) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
setMidToLo:
	mid = lo
moveUpAndIns:
	hi = len(newSlice) - 2
	for hi >= mid {
		move(newSlice, hi, hi+1)
		hi -= 1
	}
	newSlice[mid] = val
	return newSlice, mid
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
