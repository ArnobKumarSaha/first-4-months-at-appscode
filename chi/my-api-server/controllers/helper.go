package controllers

import "strconv"

func isEqual(a interface{},b interface{}) bool{
	var A, B []byte

	switch u := a.(type) {
	case int:
		A = []byte(strconv.Itoa(u))
	case string:
		A = []byte(u)
	case int64:
		A = []byte(strconv.FormatInt(int64(u),10))
	}

	switch u := b.(type) {
	case int:
		B = []byte(strconv.Itoa(u))
	case string:
		B = []byte(u)
	case int64:
		B = []byte(strconv.FormatInt(int64(u),10))
	}

	for len(A) != len(B){
		return false
	}
	for i,j := range A{
		if B[i] != j{
			return false
		}
	}
	return true
}
