package go_effect_sort

import "cmp"

type Empty = struct{}

type LessThan[T any] func(slice []T, aIdx int, bIdx int) bool
type LessThanOrEqual[T any] func(slice []T, aIdx int, bIdx int) bool
type GreaterThan[T any] func(slice []T, aIdx int, bIdx int) bool
type GreaterThanOrEqual[T any] func(slice []T, aIdx int, bIdx int) bool
type EqualOrder[T any] func(slice []T, aIdx int, bIdx int) bool
type UnequalOrder[T any] func(slice []T, aIdx int, bIdx int) bool
type EqualValue[T any] func(slice []T, aIdx int, bIdx int) bool
type UnequalValue[T any] func(slice []T, aIdx int, bIdx int) bool
type LessThanExternal[T any] func(slice []T, idx int, val T) bool
type LessThanOrEqualExternal[T any] func(slice []T, idx int, val T) bool
type GreaterThanExternal[T any] func(slice []T, idx int, val T) bool
type GreaterThanOrEqualExternal[T any] func(slice []T, idx int, val T) bool
type EqualOrderExternal[T any] func(slice []T, idx int, val T) bool
type UnequalOrderExternal[T any] func(slice []T, idx int, val T) bool
type EqualValueExternal[T any] func(slice []T, idx int, val T) bool
type UnequalValueExternal[T any] func(slice []T, idx int, val T) bool
type Swap[T any] func(slice []T, aIdx int, bIdx int)
type Move[T any] func(slice []T, old int, new int)
type Remove[T any, X any] func(slice []T, idx int) (val T, effectVals X)
type Insert[T any, X any] func(slice []T, idx int, val T, effectVals X)

func IsSorted[T any](slice []T, greaterThan GreaterThan[T]) bool {
	var i int = 0
	var ii int = 1
	var n int = len(slice)
	for ii < n {
		if greaterThan(slice, i, ii) {
			return false
		}
		i += 1
		ii += 1
	}
	return true
}

func LessThanOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] < slice[bIdx]
}
func LessThanOrEqualOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] <= slice[bIdx]
}
func GreaterThanOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] > slice[bIdx]
}
func GreaterThanOrEqualOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] >= slice[bIdx]
}
func EqualOrderOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] == slice[bIdx]
}
func UnequalOrderOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] != slice[bIdx]
}
func EqualValueOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] == slice[bIdx]
}
func UnequalValueOrd[T cmp.Ordered](slice []T, aIdx int, bIdx int) bool {
	return slice[aIdx] != slice[bIdx]
}
func LessThanOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] < val
}
func LessThanOrEqualOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] <= val
}
func GreaterThanOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] > val
}
func GreaterThanOrEqualOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] >= val
}
func EqualOrderOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] == val
}
func UnequalOrderOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] != val
}
func EqualValueOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] == val
}
func UnequalValueOrdExt[T cmp.Ordered](slice []T, idx int, val T) bool {
	return slice[idx] != val
}
func SwapNoSideEffect[T any](slice []T, aIdx int, bIdx int) {
	tmp := slice[aIdx]
	slice[aIdx] = slice[bIdx]
	slice[bIdx] = tmp
}
func MoveNoSideEffect[T any](slice []T, old int, new int) {
	slice[new] = slice[old]
}
func RemoveNoSideEffect[T any](slice []T, idx int) (val T, effectVals Empty) {
	return slice[idx], Empty{}
}
func InsertNoSideEffect[T any](slice []T, idx int, val T, effectVals Empty) {
	slice[idx] = val
}
