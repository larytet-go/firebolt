package merge

import (
	"testing"
)

func idxExists(A, B []int, aIdx, bIdx int) bool {
	return (len(A) > aIdx) && (len(B) > bIdx)
}

func copy_the_rest(src []int, srcIdx int, dst []int, dstIdx int) {
	if srcIdx >= len(src) {
		return
	}
	copy(dst[dstIdx:], src[srcIdx:])
}

/*
Inpit Two arrays [int] sorted
Outuput merge of two input arrays sorted
*/
func mergeSortedArrays(A []int, B []int) (C []int) {
	C = make([]int, len(A)+len(B))

	var aIdx, bIdx, cIdx int
	for {
		if idxExists(A, B, aIdx, bIdx) && (A[aIdx] > B[bIdx]) {
			C[cIdx] = B[bIdx]
			bIdx += 1
			cIdx += 1
		}

		if idxExists(A, B, aIdx, bIdx) && (A[aIdx] <= B[bIdx]) {
			C[cIdx] = A[aIdx]
			aIdx += 1
			cIdx += 1
		}

		if len(B) <= bIdx {
			copy_the_rest(A, aIdx, C, cIdx)
			aIdx = len(A)
			return
		}

		if len(A) <= aIdx {
			copy_the_rest(B, bIdx, C, cIdx)
			bIdx = len(B)
			return
		}

		if len(A) <= aIdx && len(B) <= bIdx {
			return
		}
	}
}

func compareArrays(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	for i := range A {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	res := mergeSortedArrays([]int{1, 2, 3}, []int{1, 3, 4})

	if !compareArrays(res, []int{1, 1, 2, 3, 3, 4}) {
		t.Errorf("got %v", res)
	}
}
