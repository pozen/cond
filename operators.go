package cond

type Operator string

// Operator func define
// params：
//		l   left value
//		op  Operator
//		r   right value
//		reg register(runtime context, contains keys & values which uesed to compare with the Cond)
// return：
//		bool (result of condition check)
type OpFunc func(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool

// operator funcs map
var _operators map[Operator]OpFunc

func AddOperatorFunc(op Operator, f OpFunc) {
	_operators[op] = f
}

func init() {
	_operators = map[Operator]OpFunc{}
	AddOperatorFunc(opMain, opMainFunc)
	AddOperatorFunc(opAnd, opAndFunc)
	AddOperatorFunc(opOr, opOrFunc)
	AddOperatorFunc(opRegex, opRegexFunc)
	AddOperatorFunc(opContain, opContainFunc)
	AddOperatorFunc(opIn, opInfunc)
	AddOperatorFunc(opNin, opNinfunc)
	AddOperatorFunc(opEQ, opCompareFunc)
	AddOperatorFunc(opNE, opCompareFunc)
	AddOperatorFunc(opGT, opCompareFunc)
	AddOperatorFunc(opGTE, opCompareFunc)
	AddOperatorFunc(opLT, opCompareFunc)
	AddOperatorFunc(opLTE, opCompareFunc)
}
