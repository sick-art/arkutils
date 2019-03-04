package sliceutil_test

import (
	"github.com/sick-art/arkutils/sliceutil"
	"testing"
)

func ok(t *testing.T,exp interface{},result interface{}){
	if(result==exp){
		t.Log("Test successful")
	} else {
		t.Errorf("Test Failed. Expected %v but got %v",exp, result)
	}
}

func handleError(t *testing.T, err interface{}){
	if(err!=nil){
		t.Errorf("Error occured %s",err)
	}
}

func TestFindValueInSlice(t *testing.T){
	t.Log("Find value in slice")
	//Test setup
	var a []int
	a = make([]int, 2)
	a[0]=9
	a[1]=2
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
	a[0]="a"
	result,err := sliceutil.IsNotEmpty(a)
	ok(t,true,result)
	handleError(t,err)
}