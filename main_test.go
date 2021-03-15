package main

import "testing"

func verifyCombinationResult(result map[int][][]int, t *testing.T) bool {
	if result == nil || len(result) < 1 {
		return false
	}
	failFlag := false
	for k, v := range result {
		if v == nil {
			t.Logf("sum [%v] result contains no combination list", k)
			failFlag = true
			continue
		}
		for _, aCombo := range v {
			tr := Sum(aCombo)
			if tr != k {
				t.Logf("sum [%v] result contains an invalid combination [%v]", k, aCombo)
				failFlag = true
			}
		}
	}
	return failFlag
}

func TestNoInputValue(t *testing.T) {
	GenerateComboSums(nil, 2)
	GenerateComboSums([]int{}, 2)
}

func TestUnSortedArray(t *testing.T) {
	result := GenerateComboSums([]int{5, 4, 3, 2, 1, 7, 8, 9}, 2)
	verifyCombinationResult(result, t)
}

func TestSortedWithSmallAndLargeNumbers(t *testing.T) {
	result := GenerateComboSums([]int{5, 4, 3, 2, 1, 7, 100, 1010, 1107}, 2)
	verifyCombinationResult(result, t)
}

// none of the combinations sum to 10
func TestBadResult(t *testing.T) {
	badResult := map[int][][]int{}
	sumList := make([][]int, 2)
	sumList[0] = []int{1, 2}
	sumList[1] = []int{3, 4}
	badResult[10] = sumList
	failFlag := verifyCombinationResult(badResult, t)
	if !failFlag {
		t.Fail()
	}
}

// all combinations sum to 7
func TestGoodResult(t *testing.T) {
	badResult := map[int][][]int{}
	sumList := make([][]int, 2)
	sumList[0] = []int{1, 6}
	sumList[1] = []int{3, 4}
	badResult[7] = sumList
	failFlag := verifyCombinationResult(badResult, t)
	if failFlag {
		t.Fail()
	}
}

// some combinations sum to 7
func TestBadResultMost7(t *testing.T) {
	badResult := map[int][][]int{}
	sumList := make([][]int, 3)
	sumList[0] = []int{1, 6}
	sumList[1] = []int{1, 8}
	sumList[2] = []int{3, 4}
	badResult[7] = sumList
	failFlag := verifyCombinationResult(badResult, t)
	if !failFlag {
		t.Fail()
	}
}
