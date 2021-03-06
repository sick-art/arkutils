package sliceutil_test

import (
	"github.com/sick-art/arkutils/sliceutil"
	"testing"
)

func ok(t *testing.T,exp interface{},result interface{}){
	if( result==exp ){
		t.Log("Test successful")
	} else {
		t.Errorf("Test Failed. Expected %v but got %v",exp, result)
	}
}

func handleError(t *testing.T, err interface{}){
	if( err!=nil ){
		t.Errorf("Error occured %s",err)
	}
}

func TestFindValueInSlice(t *testing.T){
	t.Log("Find value in slice")
	//Test setup
	var a []int
	a = make([]int, 2)
	a[0] = 9
	a[1] = 2
	//Test call
	result := sliceutil.Find(a, func(i interface{}) bool {return i==2 })
	//Test
	ok(t,1,result)
}

func TestContainsValueInSlice(t *testing.T){
	t.Log("contains value in slice")
	//Test setup
	var a []string
	a = make([]string,3)
	a[0] = "aaabc"
	a[1] = "arkutils"
	a[2] = "cabac"
	var findString string = "arkutils"

	//Test call
	result,err := sliceutil.Contains(a, findString)
	
	//Test
	ok(t,true,result)
	handleError(t,err)
}

func TestIsEmpty(t *testing.T){
	var a []string
	a = make([]string,0)
	result,err := sliceutil.IsEmpty(a)
	ok(t,true,result)
	handleError(t,err)
}

func TestIsNotEmpty(t *testing.T){
	var a []string
	a = make([]string,1)
	a[0] = "a"
	result,err := sliceutil.IsNotEmpty(a)
	ok(t,true,result)
	handleError(t,err)
}

func TestIsEqual(t *testing.T){
	var a[]int
	a = make([]int,2)
	a[0] = 9
	a[1] = 2
	var b[]int
	b = make([]int,2)
	b[0] = 9
	b[1] = 2
	
	result := sliceutil.IsEqual(a,b)

	ok(t,true,result)
}

func TestIsSameLength(t *testing.T){
	var a[]int
	a = make([]int,2)
	a[0] = 9
	a[1] = 2
	var b[]int
	b = make([]int,2)
	b[0] = 9
	b[1] = 2
	
	result := sliceutil.IsSameLength(a,b)

	ok(t,true,result)
}

func TestIsSameType(t *testing.T){
	var a[]int
	a = make([]int,2)
	a[0] = 9
	a[1] = 2
	var b[]int
	b = make([]int,2)
	b[0] = 9
	b[1] = 2
	
	result := sliceutil.IsSameType(a,b)

	ok(t,true,result)
}

func TestLastIndexOf(t *testing.T){
	var a[]int
	a = make([]int,5)
	a[0] = 9
	a[1] = 2
	a[2] = 7
	a[3] = 2
	a[4] = 10
	
	result := sliceutil.LastIndexOf(a,2,5)

	ok(t,3,result)
}

func TestRemoveElement(t *testing.T){
	var a[]string
	a = make([]string,3)
	a[0] = "a"
	a[1] = "abac"
	a[2] = "aabc"
	
	result := sliceutil.RemoveElement(a,1)
	
	var b[]string
	b = make([]string,2)
	b[0] = "a"
	b[1] = "aabc"

	testResult:= sliceutil.IsEqual(result,b)
	
	ok(t,true,testResult)
}
