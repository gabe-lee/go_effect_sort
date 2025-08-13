package go_effect_sort

import (
	"slices"
	"testing"
)

func FuzzBinaryInsert(f *testing.F) {
	var nilSlice []byte
	var emptySlice []byte = make([]byte, 0, 10)
	f.Add(nilSlice, byte(5))
	f.Add(emptySlice, byte(5))
	f.Add([]byte{0, 1, 2, 3, 4}, byte(5))
	f.Add([]byte{1, 2, 3, 4}, byte(5))
	f.Add([]byte{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}, byte(5))
	f.Add([]byte{0, 1, 2, 3, 4, 6, 7, 8, 10}, byte(5))
	f.Add([]byte{5, 5, 5, 5, 6}, byte(5))
	f.Add([]byte{5, 5, 5, 6}, byte(5))
	f.Add([]byte{6, 6, 6, 6}, byte(5))
	f.Add([]byte{6, 6, 6, 6, 6}, byte(5))
	f.Fuzz(func(t *testing.T, a []byte, b byte) {
		slices.Sort(a)
		var sum uint64 = 0
		for _, b := range a {
			sum += uint64(b) + 1
		}
		var expSum = sum + uint64(b) + 1
		var expLen = len(a) + 1
		aa, _ := BinaryInsert(a, b, EqualOrderOrdExt, GreaterThanOrdExt, MoveNoSideEffect)
		var gotSum uint64 = 0
		var gotLen = len(aa)
		if gotLen != expLen {
			t.Errorf("\ntest case failed: len mismatch\nINSERT VAL: %d\nOLD SLICE: %v\nNEW SLICE: %v\nEXP LEN: %d\nGOT LEN: %d\n", b, a, aa, expLen, gotLen)
		}
		for _, b := range aa {
			gotSum += uint64(b) + 1
		}
		if gotSum != expSum {
			t.Errorf("\ntest case failed: sum mismatch\nINSERT VAL: %d\nOLD SLICE: %v\nNEW SLICE: %v\nEXP SUM: %d\nGOT SUM: %d\n", b, a, aa, expSum, gotSum)
		}
		var ii int = 1
		var i int = 0
		for ii < gotLen {
			if aa[i] > aa[ii] {
				t.Errorf("\ntest case failed: not sorted\nINSERT VAL: %d\nOLD SLICE: %v\nNEW SLICE: %v\nVAL : %d > %d\nIDX: %d < %d\n", b, a, aa, aa[i], aa[ii], i, ii)
			}
			i += 1
			ii += 1
		}
	})
}

func FuzzBinarySearch(f *testing.F) {
	var nilSlice []byte
	var emptySlice []byte = make([]byte, 0, 10)
	f.Add(nilSlice, byte(5))
	f.Add(emptySlice, byte(5))
	f.Add([]byte{0, 1, 2, 3, 4}, byte(5))
	f.Add([]byte{1, 2, 3, 4}, byte(5))
	f.Add([]byte{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}, byte(5))
	f.Add([]byte{0, 1, 2, 3, 4, 6, 7, 8, 10}, byte(5))
	f.Add([]byte{5, 5, 5, 5, 6}, byte(5))
	f.Add([]byte{5, 5, 5, 6}, byte(5))
	f.Add([]byte{6, 6, 6, 6}, byte(5))
	f.Add([]byte{6, 6, 6, 6, 6}, byte(5))
	f.Fuzz(func(t *testing.T, a []byte, b byte) {
		slices.Sort(a)
		var minValidIdx int
		var maxValidIdx int
		var existsInList bool
		for i, bb := range a {
			if !existsInList {
				if bb == b {
					existsInList = true
					minValidIdx = i
					maxValidIdx = i
				}
			} else {
				if bb > b {
					maxValidIdx = i - 1
					break
				} else {
					maxValidIdx = i
				}
			}
		}
		foundIdx, found := BinarySearch(a, b, EqualValueOrdExt, GreaterThanOrdExt)
		if found && !existsInList {
			t.Errorf("\ntest case failed: value does not exist in list but was 'found' by BinarySearch\nSEARCH VAL: %d\nSLICE: %v\nBAD FOUND IDX: %d\n", b, a, foundIdx)
		}
		if !found && existsInList {
			t.Errorf("\ntest case failed: value exists in list but was not 'found' by BinarySearch\nSEARCH VAL: %d\nSLICE: %v\n", b, a)
		}
		if foundIdx < minValidIdx || foundIdx > maxValidIdx {
			t.Errorf("\ntest case failed: found value idx outside valid range of matching bytes\nSEARCH VAL: %d\nSLICE: %v\nMIN VALID IDX: %d\nMAX VALID IDX: %d\n'FOUND' IDX: %d", b, a, minValidIdx, maxValidIdx, foundIdx)
		}
	})
}
