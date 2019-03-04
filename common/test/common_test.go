package test

import (
	"testing"

	"github.com/sick-art/arkutils/common/math/util"
	"github.com/sick-art/arkutils/common/text/similarity"
)

func ok(t *testing.T, err error) {
	//To be filled
}

func TestLongestCommonSubsequence(t *testing.T) {
	t.Log("Test")
	var a, b string = "New York", "New Hampshire"
	result := similarity.LongestCommonSubsequence(a, b)
	t.Log(result)
	if len(result) == 5 {
		t.Log("Yay! We passed here.")
	} else {
		t.Errorf("Nah! We can do better.")
	}
}


func TestMaxMin(t *testing.T) {
	t.Log("Test")
	result1 := util.Max(234, 453)
	if result1 == 453 {
		t.Log("Yes we knew Max of (234, 453). 453! That's correct.")
	} else {
		t.Log("Don't know yet!")
	}
	result2 := util.Min(234, 453)
	if result2 == 234 {
		t.Log("Yes we knew Min of (234, 453). 234! That's correct.")
	} else {
		t.Log("Don't know yet!")
	}
}
