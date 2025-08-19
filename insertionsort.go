package go_effect_sort

import (
	ll "github.com/gabe-lee/go_list_like"
)

func InsertionSort[T any](slice ll.ListLike[T], greaterThan func(slice ll.SliceLike[T], idx int, val T) bool, remove func(slice ll.ListLike[T], idx int) (val T), move func(slice ll.SliceLike[T], old int, new int), insert func(slice ll.ListLike[T], idx int, val T)) {
	var n int = slice.Len()
	var i int = 1
	var j int
	var jj int
	var elem T
	for i < n {
		elem = remove(slice, i)
		j = i - 1
		jj = i
		for j >= 0 && greaterThan(slice, j, elem) {
			move(slice, j, jj)
			j -= 1
			jj -= 1
		}
		insert(slice, jj, elem)
		i += 1
	}
}
