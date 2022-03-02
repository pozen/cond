package cond

import (
	"reflect"
)

const opContain Operator = "$contain"

func opContainFunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {
	lVal := s.loadValFromSession(l, reg)
	if lVal == nil {
		return false
	}

	//rType := reflect.TypeOf(r)
	lType := reflect.TypeOf(lVal)

	if lType.Kind() != reflect.Slice {
		panic("$contain left value must be a slice")
	}

	lElementType := lType.Elem()

	if lElementType.Kind() == reflect.Interface {
		panic("$contain: type of left value can not be []interface{}")
	}

	//if getValType(rType.Kind()) != getValType(lElementType.Kind()) {
	//	panic("$contain: type mismatch between left value & right value")
	//}

	lValue := reflect.ValueOf(lVal)

	for i := 0; i < lValue.Len(); i++ {
		lrIndex := lValue.Index(i).Interface()

		if compareValue(s, opEQ, r, lrIndex, reg) {
			return true
		}
	}

	return false
}
