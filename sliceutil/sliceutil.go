package sliceutil

import (
	"fmt"
	"reflect"
	"errors"
)

/*
	Usage: find(a, func(i interface{}) bool { return i == 5 })
	returns index of a value in a given slice
	Accepts:
		- slice : slice of an array
		  f     : filter function for finding the element
*/
func Find(slice interface{}, f func(interface{}) bool) int {
    switch reflect.TypeOf(slice).Kind() {
    case reflect.Slice:
        values := reflect.Indirect(reflect.ValueOf(slice))
        for i := 0; i < values.Len(); i++ {
            if f(values.Index(i).Interface()) {
                return i
            }
        }
    }
    return -1
}

/*
	Usage: contains(a, 34)
	returns index of a value in a given slice
	Accepts:
		- slice : slice of an array
		  value     : value of matching type of element in slice
*/
func Contains(slice interface{}, value interface{}) (bool,error) {
	empty,err_empty := IsEmpty(slice)
	isNil,err_isNil := IsNil(slice)
	if( empty==true) {
		return false, errors.New(fmt.Sprintf("Slice can't be empty %s",err_empty))
	}
	if( isNil==true) {
		return false, errors.New(fmt.Sprintf("Slice can't be nil %s",err_isNil))
	}

	var index int = Find(slice, func(i interface{}) bool { return i==value })
	if(index>-1) {
		return true, nil
	} else {
		return false, nil
	}
}

/*
	Usage: IsNil(a)
	Returns a boolean value of IsNil of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsNil(slice interface{}) (bool,error) {
	return (slice==nil),nil;
}

/*
	Usage: IsNotNil(a)
	Returns a boolean value of IsNotNil of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsNotNil(slice interface{}) (bool,error) {
	return (!(slice==nil)),nil;
}

/*
	Usage: IsEmpty(a)
	Returns a boolean value of IsEmpty of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsEmpty(slice interface{}) (bool,error) {
	if(slice==nil) {
		return false, errors.New("Slice can't be nil")
	}
	return (reflect.ValueOf(slice).Len()==0),nil
}


/*
	Usage: IsNotEmpty(a)
	Returns a boolean value of IsNotEmpty of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsNotEmpty(slice interface{}) (bool,error) {
	return (!(reflect.ValueOf(slice).Len()==0)),nil;
}


/*
	Usage: IsEqual(a)
	Returns a boolean value of IsEqual for two slices
	Accepts:
		- slice  : slice of an array
		  anotherSlice : slice of an array

*/
func IsEqual(slice interface{}, anotherSlice interface{}) (bool) {
	return reflect.DeepEqual(slice,anotherSlice)
}

/*
	Usage: IsSameLength(a)
	Returns a boolean value of IsSameLength for two slices
	Accepts:
		- slice        : slice of an array
		  anotherSlice : slice of an array
*/
func IsSameLength(slice interface{}, anotherSlice interface{}) (bool) {
	return reflect.ValueOf(slice).Len()==reflect.ValueOf(anotherSlice).Len()
}

/*
	Usage: IsSameType(a)
	Returns a boolean value of IsSameType for two slices
	Accepts:
		- slice        : slice of an array
		  anotherSlice : slice of an array
*/
func IsSameType(slice interface{}, anotherSlice interface{}) (bool) {
	return reflect.TypeOf(slice)==reflect.TypeOf(anotherSlice)
}

func LastIndexOf(slice interface{}, valueToFind interface{}, startIndex int) (int) {
	if(slice==nil){
		return -1
	}
	if( startIndex < 0 ){
		return -1
	} else if( startIndex >= reflect.ValueOf(slice).Len()){
		startIndex = reflect.ValueOf(slice).Len()-1
	}
	if(valueToFind==nil){
		
	}
}
