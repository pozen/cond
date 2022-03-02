package cond

import (
	"reflect"

	"github.com/shopspring/decimal"
)

const (
	opEQ  Operator = "$eq"
	opGT           = "$gt"
	opGTE          = "$gte"
	opLT           = "$lt"
	opLTE          = "$lte"
	opNE           = "$ne"
)

var compareOperators []Operator = []Operator{opEQ, opGT, opGTE, opLT, opLTE, opNE}

func isCompareOperator(op Operator) bool {
	for _, v := range compareOperators {
		if op == v {
			return true
		}
	}
	return false
}

func opCompareFunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {
	if !isCompareOperator(op) {
		panic("not a compare operator :" + string(op))
	}

	lval := s.loadValFromSession(l, reg)
	if lval == nil {
		return false
	}

	switch reflect.TypeOf(lval).Kind() {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Bool, reflect.String:
		return compareValue(s, op, lval, r, reg)
	default:
		panic("not comparable type : " + reflect.TypeOf(lval).Kind().String())
	}
}

type valueType int

const (
	valueTypeBool valueType = iota
	valueTypeDecimal
	valueTypeString
	valueTypeOther
)

func getValType(kind reflect.Kind) valueType {
	switch kind {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return valueTypeDecimal
	case reflect.Bool:
		return valueTypeBool
	case reflect.String:
		return valueTypeString
	default:
		return valueTypeOther
	}
}

func getDecimalValue(i interface{}) decimal.Decimal {
	switch v := i.(type) {
	case int:
		return decimal.New(int64(v), 0)
	case int8:
		return decimal.New(int64(v), 0)
	case int16:
		return decimal.New(int64(v), 0)
	case int32:
		return decimal.New(int64(v), 0)
	case int64:
		return decimal.New(int64(v), 0)
	case uint:
		return decimal.New(int64(v), 0)
	case uint8:
		return decimal.New(int64(v), 0)
	case uint16:
		return decimal.New(int64(v), 0)
	case uint32:
		return decimal.New(int64(v), 0)
	case uint64:
		return decimal.New(int64(v), 0)
	case float32:
		return decimal.NewFromFloat32(v)
	case float64:
		return decimal.NewFromFloat(v)
	}
	return decimal.New(0, 0)
}

func compareValue(s *State, op Operator, i1 interface{}, ir interface{}, reg map[string]interface{}) bool {

	i2 := s.loadVarible(ir, reg)

	if getValType(reflect.TypeOf(i1).Kind()) != getValType(reflect.TypeOf(i2).Kind()) {
		return false
	}

	type1 := getValType(reflect.TypeOf(i1).Kind())
	switch type1 {
	case valueTypeBool:
		val1, val2 := i1.(bool), i2.(bool)
		if op == opEQ {
			return val1 == val2
		} else if op == opNE {
			return val1 != val2
		} else {
			return false
		}
	case valueTypeDecimal:
		val1, val2 := getDecimalValue(i1), getDecimalValue(i2)
		if op == opEQ {
			return val1.Equal(val2)
		} else if op == opNE {
			return !val1.Equal(val2)
		} else if op == opGT {
			return val1.GreaterThan(val2)
		} else if op == opGTE {
			return val1.GreaterThanOrEqual(val2)
		} else if op == opLT {
			return val1.LessThan(val2)
		} else if op == opLTE {
			return val1.LessThanOrEqual(val2)
		} else {
			return false
		}
	case valueTypeString:
		val1, val2 := i1.(string), i2.(string)
		if op == opEQ {
			return val1 == val2
		} else if op == opNE {
			return val1 != val2
		} else if op == opGT {
			return val1 > val2
		} else if op == opGTE {
			return val1 >= val2
		} else if op == opLT {
			return val1 < val2
		} else if op == opLTE {
			return val1 <= val2
		} else {
			return false
		}
	default:
		return false
	}
}
