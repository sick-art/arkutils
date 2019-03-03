package sliceutil

import (
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
	if(IsEmpty(slice) || isNil(slice)) {
		return nil, errors.New("Slice can't be empty")
	}
	var index int = Find(slice, func(i interface{}) bool { return i==value })
	if(index>-1) {
		return true
	} else {
		return false
	}
}

/*
	Usage: IsNil(a)
	Returns a boolean value of IsNil of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsNil(slice interface{}) (bool,error) {
	return slice==nil;
}

/*
	Usage: IsNotNil(a)
	Returns a boolean value of IsNotNil of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsNotNil(slice interface{}) (bool,error) {
	var isNil, error = IsNil(slice);
	return !isNil;
}

/*
	Usage: IsEmpty(a)
	Returns a boolean value of IsEmpty of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsEmpty(slice interface{}) (bool,error) {
	if(slice==nil) {
		return errors.New("Slice passed is nil.")
	}
	return len(slice)==0
}


/*
	Usage: IsNotEmpty(a)
	Returns a boolean value of IsNotEmpty of a given slice
	Accepts:
		- slice  : slice of an array

*/
func IsNotEmpty(slice interface{}) (bool,error) {
	var isEmpty,error = IsEmpty(slice);
	return !isEmpty;
}