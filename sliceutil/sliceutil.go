package sliceutil

import (
	"reflect"
)

/*
	Usage: find(a, func(i interface{}) bool { return i == 5 })
	finds index of a value in a given slice
	Accepts:
		- slice : a generic slice of an array
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