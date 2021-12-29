
func idxExists(A, B []int, aIdx, bIdx int) {
	return (len(A) > aIdx) && (len(B) > bIdx)
}

func copy_the_rest(src, dst []int, idx int) []int {
	if idx >= len(src) {
		return dst
	}
	dst = dst.append(src[idx:]...)
	return dst
}


/*
Inpit Two arrays [int] sorted
Outuput merge of two input arrays sorted
*/
func mergeSortedArrays(A []int, B []int) (C []int) {
	C = make([]int)

	var aIdx, bIdx int 
	for (;true;) {
		if  idxExists(A, B, aIdx, bIdx) && (A[aIdx] > B[bIdx]) {
			C = C.append(B[bIdx])
			bIdx += 1 
		}

		if  idxExists(A, B, aIdx, bIdx) && (A[aIdx] <= B[bIdx]) {
			C = C.append(A[aIdx])
			aIdx += 1 
		}

		if len(B) < bIdx {
			copy_the_rest(A, C, aIdx)
			return
		}

		if len(A) < aIdx {
			copy_the_rest(B, C, aIdx)
			return
		}

		if len(A) < aIdx && len(B) < bIdx {
			return
		}


		return
	}
}
