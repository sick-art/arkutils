package similarity

import (
	"fmt"
	"strings"
	"github.com/sick-art/arkutils/common/math/util"
	"github.com/sick-art/arkutils/stringutil"
)

func LongestCommonSubsequence(left, right string) string {
	fmt.Println(left, right)
	if stringutil.IsEmpty(left) || stringutil.IsEmpty(right) {
		panic("Inputs must not be null")
	}
	leftLength := len(left)
	rightLength := len(right)
	length := util.Max(leftLength, rightLength)
	longestCommonSubstringArray := make([]string, length)
	lcsLengthArray := LongestCommonSubsequenceLengthArray(left, right)
	var i, j, k int = leftLength - 1, rightLength - 1, lcsLengthArray[leftLength][rightLength] - 1

	for ok := true; ok; ok = (k >= 0) {
		if left[i] == right[j] {
			longestCommonSubstringArray = append(longestCommonSubstringArray, string(left[i]))
			i = i - 1
			j = j - 1
			k = k - 1
		} else if lcsLengthArray[i+1][j] < lcsLengthArray[i][j+1] {
			i = i - 1
		} else {
			j = j - 1
		}
	}
	return stringutil.Reverse(strings.Join(longestCommonSubstringArray,""))
}

/* Computes the lcsLengthArray for the sake of doing the actual LCS calculation.
This is the Dynamic Programming portion of the algorith, and tis the reason for runtime complexity O(m*n),
where m = len(left) and n = len(right)
*/
func LongestCommonSubsequenceLengthArray(left, right string) [][]int {

	// initialize 2D array
	lcsLengthArray := make([][]int, len(left)+1)
	for i := range lcsLengthArray {
		lcsLengthArray[i] = make([]int, len(right)+1)
	}

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if i == 0 {
				lcsLengthArray[i][j] = 0
			}
			if j == 0 {
				lcsLengthArray[i][j] = 0
			}
			if left[i] == right[i] {
				lcsLengthArray[i+1][j+1] = lcsLengthArray[i][j] + 1
			} else {
				lcsLengthArray[i+1][j+1] = util.Max(lcsLengthArray[i+1][j], lcsLengthArray[i][j+1])
			}
		}
	}

	return lcsLengthArray
}
