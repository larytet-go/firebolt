package merge

import (
	"testing"
)

/*
Input: Two arrays []int sorted
Outuput: merge of two input arrays sorted
*/
func mergeSortedArrays(A []int, B []int) (C []int) {
	C = make([]int, len(A)+len(B))
	tempC := C[:]

	for {
		if len(B) > 0 && A[0] > B[0] {
			tempC[0] = B[0]
			B = B[1:]
			tempC = tempC[1:]
		}

		if len(A) > 0 && A[0] <= B[0] {
			tempC[0] = A[0]
			A = A[1:]
			tempC = tempC[1:]
		}

		if len(A) == 0 && len(B) > 0 {
			copy(tempC, B)
			return
		}

		if len(B) == 0 && len(A) > 0 {
			copy(tempC, A)
			return
		}

		if len(A) == 0 && len(B) == 0 {
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
	res := mergeSortedArrays([]int{1, 2, 3}, []int{1, 3, 4, 4, 5})

	expected := []int{1, 1, 2, 3, 3, 4, 4, 5}
	if !compareArrays(res, expected) {
		t.Errorf("got/expected:\n%v\n%v", res, expected)
	}
}
