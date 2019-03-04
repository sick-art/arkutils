package stringutil


func IsEmpty(s string) bool {
	return len(s)==0
}


/*
Basic string reversal function
*/
func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}
  
