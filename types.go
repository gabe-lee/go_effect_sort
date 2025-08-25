package go_effect_sort

import (
	"cmp"

	ll "github.com/gabe-lee/go_list_like"
)

// Convenience wrapper around `github.com/gabe-lee/go_slice_like.New()`
func NewSliceLike[T any](slicePtr *[]T) ll.SliceAdapterIndirect[T] {
	return ll.NewSliceAdapterIndirect(slicePtr)
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
	return slice.Get(aIdx) < slice.Get(bIdx)
}
func LessThanOrEqualOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) <= slice.Get(bIdx)
}
func GreaterThanOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) > slice.Get(bIdx)
}
func GreaterThanOrEqualOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) >= slice.Get(bIdx)
}
func EqualOrderOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) == slice.Get(bIdx)
}
func UnequalOrderOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) != slice.Get(bIdx)
}
func EqualValueOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) == slice.Get(bIdx)
}
func UnequalValueOrd[T cmp.Ordered](slice ll.SliceLike[T], aIdx int, bIdx int) bool {
	return slice.Get(aIdx) != slice.Get(bIdx)
}

func LessThanOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) < slice.Get(int(idxMap.Get(bIdx)))
}
func LessThanOrEqualOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) <= slice.Get(int(idxMap.Get(bIdx)))
}
func GreaterThanOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) > slice.Get(int(idxMap.Get(bIdx)))
}
func GreaterThanOrEqualOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) >= slice.Get(int(idxMap.Get(bIdx)))
}
func EqualOrderOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) == slice.Get(int(idxMap.Get(bIdx)))
}
func UnequalOrderOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) != slice.Get(int(idxMap.Get(bIdx)))
}
func EqualValueOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) == slice.Get(int(idxMap.Get(bIdx)))
}
func UnequalValueOrdMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], aIdx int, bIdx int) bool {
	return slice.Get(int(idxMap.Get(aIdx))) != slice.Get(int(idxMap.Get(bIdx)))
}

func LessThanOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) < val
}
func LessThanOrEqualOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) <= val
}
func GreaterThanOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) > val
}
func GreaterThanOrEqualOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) >= val
}
func EqualOrderOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) == val
}
func UnequalOrderOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) != val
}
func EqualValueOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) == val
}
func UnequalValueOrdExt[T cmp.Ordered](slice ll.SliceLike[T], idx int, val T) bool {
	return slice.Get(idx) != val
}

func LessThanOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) < val
}
func LessThanOrEqualOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) <= val
}
func GreaterThanOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) > val
}
func GreaterThanOrEqualOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) >= val
}
func EqualOrderOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) == val
}
func UnequalOrderOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) != val
}
func EqualValueOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) == val
}
func UnequalValueOrdExtMapped[T cmp.Ordered, I Index](slice ll.SliceLike[T], idxMap ll.SliceLike[I], idx int, val T) bool {
	return slice.Get(int(idxMap.Get(idx))) != val
}

func SwapNoSideEffect[T any](slice ll.SliceLike[T], aIdx int, bIdx int) {
	ll.Swap(slice, aIdx, bIdx)
}
func MoveNoSideEffect[T any](slice ll.SliceLike[T], old int, new int) {
	ll.Move(slice, old, new)
}
func RemoveNoSideEffect[T any](slice ll.ListLike[T], idx int) (val T) {
	return slice.Get(idx)
}
func InsertNoSideEffect[T any](slice ll.ListLike[T], idx int, val T) {
	slice.Set(idx, val)
}
