package cond

import (
	"reflect"
)

const opMain Operator = ""

// entry func of condtion check.
func opMainFunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {
	if l == "" && op == opMain {
		condMap, ok := r.(Cond)
		if !ok {
			condMap, ok = r.(map[string]interface{})
			if !ok {
				return false
			}
		}
		for k, v := range condMap {

			// check operators first
			f, ok := _operators[Operator(k)]
			if ok {
				if ok := f(s, k, Operator(k), v, reg); !ok {
					return false
				}
				continue
			}

			switch reflect.TypeOf(v).Kind() {

			// if v is a comparable value, $equel func will be called
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
				if !opCompareFunc(s, k, opEQ, v, reg) {
					return false
				}

			// hit $Depend Operator if right value is a map. eg:
			// { "key": {"$ne": 10}}
			case reflect.Map, reflect.TypeOf(&Cond{}).Kind():
				if !opDependFunc(s, k, opDepend, v, reg) {
					return false
				}
			}
		}
	}
	return true
}
