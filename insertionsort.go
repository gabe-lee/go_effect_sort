package go_effect_sort

func InsertionSortSwap[T any](slice []T, greaterThan GreaterThan[T], swap Swap[T]) {
	var n int = len(slice)
	var i int = 1
	var j int
	var jj int
	for i < n {
		j = i - 1
		jj = i
		for j >= 0 && greaterThan(slice, j, jj) {
			swap(slice, j, jj)
			j -= 1
			jj -= 1
		}
		i += 1
	}
}

// This may be faster than `InsertionSortSwap` if swapping two elements is a heavy operation,
// but requires additional user work to provide remove/move/insert funcs and a type for `effectVals` (X)
func InsertionSortRemoveMoveInsert[T any, X any](slice []T, greaterThan GreaterThanExternal[T], remove Remove[T, X], move Move[T], insert Insert[T, X]) {
	var n int = len(slice)
	var i int = 1
	var j int
	var jj int
	var elem T
	var elemEffects X
	for i < n {
		elem, elemEffects = remove(slice, i)
		j = i - 1
		jj = i
		for j >= 0 && greaterThan(slice, j, elem) {
			move(slice, j, jj)
			j -= 1
			jj -= 1
		}
		insert(slice, jj, elem, elemEffects)
		i += 1
	}
}
