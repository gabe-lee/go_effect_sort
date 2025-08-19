package go_effect_sort

import (
	"cmp"

	ll "github.com/gabe-lee/go_list_like"
)

// Convenience wrapper around `github.com/gabe-lee/go_slice_like.New()`
func NewSliceLike[T any](slicePtr *[]T) ll.SliceAdapter[T] {
	return ll.New(slicePtr)
}

type Index interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Empty = struct{}

func IsSorted[T any](slice ll.ListLike[T], greaterThan func(slice ll.SliceLike[T], aIdx int, bIdx int) bool) bool {
	var i int = 0
	var ii int = 1
	var n int = slice.Len()
	for ii < n {
		if greaterThan(slice, i, ii) {
			return false
		}
		i += 1
		ii += 1
	}
	return true
}

func LessThanOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) < *slice.GetPtr(bIdx)
}
func LessThanOrEqualOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) <= *slice.GetPtr(bIdx)
}
func GreaterThanOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) > *slice.GetPtr(bIdx)
}
func GreaterThanOrEqualOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) >= *slice.GetPtr(bIdx)
}
func EqualOrderOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) == *slice.GetPtr(bIdx)
}
func UnequalOrderOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) != *slice.GetPtr(bIdx)
}
func EqualValueOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) == *slice.GetPtr(bIdx)
}
func UnequalValueOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return *slice.GetPtr(aIdx) != *slice.GetPtr(bIdx)
}

func LessThanOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) < *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func LessThanOrEqualOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) <= *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func GreaterThanOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) > *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func GreaterThanOrEqualOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) >= *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func EqualOrderOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) == *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func UnequalOrderOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) != *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func EqualValueOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) == *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}
func UnequalValueOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(aIdx))) != *slice.GetPtr(int(*idxMap.GetPtr(bIdx)))
}

func LessThanOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) < val
}
func LessThanOrEqualOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) <= val
}
func GreaterThanOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) > val
}
func GreaterThanOrEqualOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) >= val
}
func EqualOrderOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) == val
}
func UnequalOrderOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) != val
}
func EqualValueOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) == val
}
func UnequalValueOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return *slice.GetPtr(idx) != val
}

func LessThanOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) < val
}
func LessThanOrEqualOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) <= val
}
func GreaterThanOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) > val
}
func GreaterThanOrEqualOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) >= val
}
func EqualOrderOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) == val
}
func UnequalOrderOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) != val
}
func EqualValueOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) == val
}
func UnequalValueOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return *slice.GetPtr(int(*idxMap.GetPtr(idx))) != val
}

func SwapNoSideEffect[T any](slice ll.SliceLike[T], aIdx int, bIdx int) {
	tmp := *slice.GetPtr(aIdx)
	*slice.GetPtr(aIdx) = *slice.GetPtr(bIdx)
	*slice.GetPtr(bIdx) = tmp
}
func MoveNoSideEffect[T any](slice ll.SliceLike[T], old int, new int) {
	*slice.GetPtr(new) = *slice.GetPtr(old)
}
func RemoveNoSideEffect[T any](slice ll.ListLike[T], idx int) (val T) {
	return *slice.GetPtr(idx)
}
func InsertNoSideEffect[T any](slice ll.ListLike[T], idx int, val T) {
	*slice.GetPtr(idx) = val
}
