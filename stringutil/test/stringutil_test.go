package test

import (
	"testing"
	"github.com/sick-art/arkutils/stringutil"
)

func ok(t *testing.T, err error) {
	//To be filled
}

func TestIfStringIsEmpty(t *testing.T) {
	t.Log("Test for empty string")
	result := stringutil.IsEmpty("")
	if result == true {
		t.Log("String Empty Test Success!")
	} else {
		t.Errorf("IsEmpty test failed!")
	}
}

func TestReverse(r *testing.T) {
	r.Log("Test")
	var a string = "aarti"
	result := stringutil.Reverse(a)
	if result == "itraa" {
		r.Log("Yay!")
	} else {
		r.Log("Nah!")
	}
}
