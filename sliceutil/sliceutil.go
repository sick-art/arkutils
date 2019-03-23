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
	returns true if value is present in slice, otherwise false
	Accepts:
		- slice 	: slice of an array
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

/*
	Usage: LastIndexOf(slice,value,startFrom)
	Returns an int value of index of the value in a given slice searched from a startIndex
	Note:
		-	If slice is nil, returns -1
		-	If startIndex is less than 0, returns -1
		-	If startIndex is greater than slice length, then startIndex is one less than array length
		-	If value is nil, returns the first nil value found in the slice
	Accepts:
		-	slice   	: slice of an array
			valueToFind	: value of slice element type
			startIndex	: int value of index to start from
*/
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
		values := reflect.Indirect(reflect.ValueOf(slice))
        for i := startIndex; i >= 0; i-- {
            if values.Index(i).Interface()==nil {
                return i
            }
        }
	} else if(reflect.TypeOf(reflect.ValueOf(slice).Index(0).Interface())==reflect.TypeOf(valueToFind)){
		values := reflect.Indirect(reflect.ValueOf(slice))
        for i := startIndex; i >= 0; i-- {
            if (valueToFind==values.Index(i).Interface()) {
                return i
            }
        }
	}
	return -1
}

/*
	Usage: removeElement(slice,index)
	Returns the slice with the element at index removed
	Accepts:
		- slice        : slice of an array of int type
		  index 	   : index of element to be removed
*/
func RemoveElement(slice interface{}, index int) (interface{}) {

	v := reflect.ValueOf(slice)

    if v.Kind() != reflect.Slice {
        fmt.Errorf("Parameter must be a slice.")
        return nil
    }

    rest := make([]interface{}, 0)


    for i := 0; i < v.Len(); i++ {
        current := reflect.ValueOf(slice).Slice(0, i)
        if v.Index(i).Interface()!=v.Index(index).Interface() {
        	
            rest = append(rest, v.Index(i).Interface())

        }
        slice = current.Interface()
    }

    elemType:= reflect.TypeOf(slice).Elem()

    switch elemType.Name() {
    case "int":
    	result := make([]int, reflect.ValueOf(rest).Len())
		t := reflect.ValueOf(rest)
		for i:=0; i <t.Len(); i++ { result[i] = t.Index(i).Interface().(int) }
		return result
	case "string":
		result := make([]string, reflect.ValueOf(rest).Len())
		t := reflect.ValueOf(rest)
		for i:=0; i <t.Len(); i++ { result[i] = t.Index(i).Interface().(string) }
		return result
	case "bool":
		result := make([]bool, reflect.ValueOf(rest).Len())
		t := reflect.ValueOf(rest)
		for i:=0; i <t.Len(); i++ { result[i] = t.Index(i).Interface().(bool) }
		return result
	case "float32":
		result := make([]float32, reflect.ValueOf(rest).Len())
		t := reflect.ValueOf(rest)
		for i:=0; i <t.Len(); i++ { result[i] = t.Index(i).Interface().(float32) }
		return result
	default:
		return nil
    }
	

    return nil
}