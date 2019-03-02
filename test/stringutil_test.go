package stringutil_test

import (
"github.com/sick-art/arkutils/stringutil"
"testing"
)

func ok(t *testing.T, err error) {
 //To be filled
}

func TestIfStringIsEmpty(t *testing.T){
   t.Log("Test for empty string")
   result := stringutil.IsEmpty("");
   if result == true {
	t.Log("String Empty Test Success!");
   } else {
	t.Errorf("IsEmpty test failed!");
   }
}

