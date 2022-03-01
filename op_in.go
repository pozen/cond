package cond

import (
	"reflect"
)

const opIn Operator = "$in"
const opNin Operator = "$nin"

func opInfunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {

	rType := reflect.TypeOf(r)
	if rType.Kind() != reflect.Slice {
		panic("$in: right value must be slice")
	}
	rElemType := rType.Elem()

	lval := s.loadValFromSession(l, reg)
	if lval == nil {
		return false
	}

	lType := reflect.TypeOf(lval)

	//hasInterfaceList := false

	if rElemType.Kind() == reflect.Interface {
		//hasInterfaceList = true
	}

	if /*!hasInterfaceList &&*/ getValType(rElemType.Kind()) != getValType(lType.Kind()) {
		panic("$in: type mismatch")
	}

	rValue := reflect.ValueOf(r)

	for i := 0; i < rValue.Len(); i++ {
		rIndex := rValue.Index(i).Interface()

		// same type list to type item, just compare value
		//if !hasInterfaceList {
		if compareValue(s, opEQ, lval, rIndex) {
			return true
		}
		//}
	}
	return false
}

func opNinfunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {

	rType := reflect.TypeOf(r)
	if rType.Kind() != reflect.Slice {
		panic("$nin: right value must be slice")
	}
	rElemType := rType.Elem()

	lval := s.loadValFromSession(l, reg)
	if lval == nil {
		return false
	}

	lType := reflect.TypeOf(lval)

	//hasInterfaceList := false

	if rElemType.Kind() == reflect.Interface {
		//hasInterfaceList = true
	}

	if /*!hasInterfaceList &&*/ getValType(rElemType.Kind()) != getValType(lType.Kind()) {
		panic("$nin: type mismatch")
	}

	rValue := reflect.ValueOf(r)

	for i := 0; i < rValue.Len(); i++ {
		rIndex := rValue.Index(i).Interface()

		// same type list to type item, just compare value
		//if !hasInterfaceList {
		if compareValue(s, opEQ, lval, rIndex) {
			return false
		}
		//}
	}
	return true
}
