package cond

import (
	"reflect"
)

const opAnd Operator = "$and"
const opOr Operator = "$or"

// all sub conds should be true. eg: { "$and": [{}, {}] }
func opAndFunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {
	if reflect.TypeOf(r).Kind() != reflect.Slice {
		panic("right value of OP $and  must be slice type")
	}
	if subCond, ok := r.([]map[string]interface{}); ok {
		for _, v := range subCond {
			if !opMainFunc(s, "", "", v, reg) {
				return false
			}
		}
	}
	return true
}

// one of subconds should be true. eg: { "$or": [{}, {}] }
func opOrFunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {
	if reflect.TypeOf(r).Kind() != reflect.Slice {
		panic("right value of OP $or  must be slice type")
	}
	if subCond, ok := r.([]map[string]interface{}); ok {
		for _, v := range subCond {
			if opMainFunc(s, "", "", v, reg) {
				return true
			}
		}
	}
	return false
}
