package go_effect_sort

import (
	"testing"

	ll "github.com/gabe-lee/go_list_like"
)

func FuzzInsertionSort(f *testing.F) {
	var nilSlice []byte
	var emptySlice []byte = make([]byte, 0, 10)
	f.Add(nilSlice)
	f.Add(emptySlice)
	f.Add([]byte{0, 1, 2, 3, 4})
	f.Add([]byte{4, 3, 2, 1, 0})
	f.Add([]byte{1, 9, 37, 223})
	f.Add([]byte{223, 37, 9, 1})
	f.Add([]byte{56, 42, 3, 77, 22, 5, 109})
	f.Fuzz(func(t *testing.T, a []byte) {
		oldLen := len(a)
		var oldSum uint64 = 0
		for _, b := range a {
			oldSum += uint64(b) + 1
		}
		aa := ll.New(&a)
		InsertionSort(aa, GreaterThanOrdExt, RemoveNoSideEffect, MoveNoSideEffect, InsertNoSideEffect)
		newLen := len(a)
		if oldLen != newLen {
			t.Errorf("\ntest case failed: len mismatch\nSLICE: %v\nEXP LEN: %d\nGOT LEN: %d\n", a, oldLen, newLen)
		}
		var newSum uint64 = 0
		for _, b := range a {
			newSum += uint64(b) + 1
		}
		if oldLen != newLen {
			t.Errorf("\ntest case failed: sum mismatch\nSLICE: %v\nEXP SUM: %d\nGOT SUM: %d\n", a, oldSum, newSum)
		}
		var ii int = 1
		var i int = 0
		for ii < oldLen {
			if a[i] > a[ii] {
				t.Errorf("\ntest case failed: not sorted\nSLICE: %v\nVAL : %d > %d\nIDX: %d < %d\n", a, a[i], a[ii], i, ii)
			}
			i += 1
			ii += 1
		}
	})
}
