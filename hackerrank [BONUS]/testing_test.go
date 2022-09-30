package hackerrank

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findLongestSubsequence(arr []int32) int32 {
	// Write your code here
	var OddA, OddB int32
	var EvenA, EvenB int32

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		if OddA == 0 || OddB == 0 {
			if OddA == 0 && arr[i]%2 == 1 {
				OddA = int32(i + 1)
			}

			if OddB == 0 && arr[len(arr)-1-i]%2 == 1 {
				OddB = int32(len(arr) - i)
			}
		}

		if EvenA == 0 || EvenB == 0 {
			if EvenA == 0 && arr[i]%2 == 0 {
				EvenA = int32(i + 1)
			}

			if EvenB == 0 && arr[len(arr)-1-i]%2 == 0 {
				EvenB = int32(len(arr) - i)
			}
		}
	}

	sumA := OddB - OddA
	sumB := EvenB - EvenA

	if OddA > 0 && OddB > 0 && EvenA > 0 && EvenB > 0 {
		if sumA > sumB {
			return sumA + 1
		}
	} else if OddA > 0 && OddB > 0 {
		return sumA + 1
	} else if EvenA > 0 && EvenB > 0 {
		return sumB + 1
	}

	return sumB + 1
}

func TestEvenDifference(t *testing.T) {
	asserts := assert.New(t)

	testcase := []struct {
		Name   string
		Input  []int32
		Output int32
	}{
		{
			Name:   "Test-01",
			Input:  []int32{2, 4, 1, 7},
			Output: 4,
		},
		{
			Name:   "Test-02",
			Input:  []int32{7, 7, 5, 6, 2, 3, 2, 4},
			Output: 6,
		},
		{
			Name:   "Test-03",
			Input:  []int32{5, 6, 2, 3, 2, 4},
			Output: 6,
		},
		{
			Name:   "Test-04",
			Input:  []int32{87, 99, 85, 50, 93},
			Output: 4,
		},
	}

	for i, val := range testcase {
		t.Run(val.Name, func(t *testing.T) {
			result := findLongestSubsequence(val.Input)
			asserts.Equal(val.Output, result, "soal ke - %d", i+1)
		})
	}
}

func getMaximumEvenSum(val []int32) int64 {
	// Write your code here
	var sum int64
	var oddA, oddB int32

	sort.Slice(val, func(i, j int) bool {
		return val[i] < val[j]
	})

	for i := 0; i < len(val); i++ {
		if val[i] > 0 && i > 0 {
			if oddA == 0 && val[i]%2 == 1 {
				oddA = val[i]
				oddB = val[i-1]
			}

			sum += int64(val[i])
		} else if val[i] > 0 {
			if oddA == 0 && val[i]%2 == 1 {
				oddA = val[i]
			}

			sum += int64(val[i])
		}
	}

	if sum%2 != 0 {
		sumA := sum - int64(oddA)
		sumB := sum + int64(oddB)

		if sumA > sumB {
			return sumA
		}

		return sumB
	}

	return sum
}

func TestDiscountTags(t *testing.T) {
	asserts := assert.New(t)

	testcase := []struct {
		Name   string
		Input  []int32
		Output int64
	}{
		{
			Name:   "Test-01",
			Input:  []int32{6, 3, 4, -1, 9, 17},
			Output: 38,
		},
		{
			Name:   "Test-02",
			Input:  []int32{-1, -2, -3, 8, 7},
			Output: 14,
		},
		{
			Name:   "Test-03",
			Input:  []int32{2, 3, 6, -5, 10, 1, 1},
			Output: 22,
		},
		{
			Name:   "Test-04",
			Input:  []int32{1, 1, 1, 1},
			Output: 4,
		},
	}

	for i, val := range testcase {
		t.Run(val.Name, func(t *testing.T) {
			result := getMaximumEvenSum(val.Input)
			asserts.Equal(val.Output, result, "soal ke - %d", i+1)
		})
	}
}

func getMaximumScore(arr []int32, k int32) int64 {
	// Write your code here
	var score int64
	var temp, max int32
	var i int32

	for i = 0; i < k; i++ {
		if arr[i] > arr[i+1] {
			temp = arr[i]
		} else {
			temp = arr[i+1]
		}

		if temp >= max {
			score += int64(temp)

			val := temp / 3
			if temp%3 > 0 {
				val += 1
			}

			max = val

			if arr[i] > arr[i+1] {
				arr[i] = val
			} else {
				arr[i+1] = val
			}
		} else {
			score += int64(max)

			val := max / 3
			if max%3 > 0 {
				val += 1
			}

			max = val

			if max < temp {
				max = temp
			}
		}
	}

	return score
}

func TestMaximumScore(t *testing.T) {
	asserts := assert.New(t)

	testcase := []struct {
		Name   string
		Input  []int32
		Length int32
		Output int64
	}{
		{
			Name:   "Test-01",
			Input:  []int32{20, 4, 3, 1, 9},
			Length: 4,
			Output: 40,
		},
		{
			Name:   "Test-02",
			Input:  []int32{4, 5, 18, 1},
			Length: 3,
			Output: 29,
		},
		{
			Name:   "Test-03",
			Input:  []int32{1, 1, 1},
			Length: 2,
			Output: 2,
		},
		{
			Name:   "Test-04",
			Input:  []int32{4, 4, 5, 18, 1},
			Length: 3,
			Output: 27,
		},
		{
			Name:   "Test-05",
			Input:  []int32{5, 20, 4, 3, 1, 9},
			Length: 4,
			Output: 34,
		},
		{
			Name:   "Test-06",
			Input:  []int32{5, 18, 8, 5, 9, 9},
			Length: 2,
			Output: 26,
		},
	}

	for i, val := range testcase {
		t.Run(val.Name, func(t *testing.T) {
			result := getMaximumScore(val.Input, val.Length)
			asserts.Equal(val.Output, result, "soal ke - %d", i+1)
		})
	}
}

func findNumOfPairs(a []int32, b []int32) int32 {
	// Write your code here
	var sum int32

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for i := 0; i < len(a); i++ {
		if a[i] > b[0] {
			sum++
			a = append(a[:i], a[i+1:]...)
			b = append(b[:0], b[0+1:]...)
			i--
		}
	}

	return sum
}

func TestPairs(t *testing.T) {
	asserts := assert.New(t)

	testcase := []struct {
		Name   string
		InputA []int32
		InputB []int32
		Output int32
	}{
		{
			Name:   "Test-01",
			InputA: []int32{1, 2, 3},
			InputB: []int32{1, 2, 1},
			Output: 2,
		},
		{
			Name:   "Test-02",
			InputA: []int32{1, 2, 3, 4, 5},
			InputB: []int32{6, 6, 1, 1, 1},
			Output: 3,
		},
		{
			Name:   "Test-03",
			InputA: []int32{2, 3, 3},
			InputB: []int32{3, 4, 5},
			Output: 0,
		},
	}

	for i, val := range testcase {
		t.Run(val.Name, func(t *testing.T) {
			result := findNumOfPairs(val.InputA, val.InputB)
			asserts.Equal(val.Output, result, "soal ke - %d", i+1)
		})
	}
}
