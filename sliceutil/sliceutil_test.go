package sliceutil_test

import (
	"github.com/sick-art/arkutils/sliceutil"
	"testing"
)

func TestFindValueInSlice(t *testing.T){
	t.Log("Find value in slice")
	var a []int
	a = make([]int, 2)
	a[0]=9
	a[1]=2
	result := sliceutil.Find(a, func(i interface{}) bool {return i==2 })
	if (result==1) {
		t.Log("Find value in slice test successful");
	} else {
		t.Errorf("Find value test Failed!!")
	}
}

func TestContainsValueInSlice(t *testing.T){
	t.Log("contains value in slice")
	var a []string
	a = make([]string,3)
	a[0] = "aaabc"
	a[1] = "arkutils"
	a[2] = "cabac"
	var findString string = "arkutils"

	result := sliceutil.Contains(a, findString)

	if(result==true) {
		t.Log("Contains value in slice test successful")
	} else {
		t.Errorf("Contains value test Failed!!")
	}
}