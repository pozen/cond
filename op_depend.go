package cond

const opDepend Operator = "$depend"

// when the type of right value is struct (map), we can not get the operator until the right value is parsed
// eg:
// {
//	"key": {"$gt":100}
// }
func opDependFunc(s *State, l string, op Operator, r interface{}, reg map[string]interface{}) bool {

	if l == "" {
		panic("$depend l value can't be empty")
	}
	rvals, ok := r.(map[string]interface{})
	if !ok {
		rvals, ok = r.(Cond)
		if !ok {
			panic("$depend right value must be map[string]interface{}")
		}
	}

	for k, v := range rvals {
		kop := Operator(k)
		switch kop {
		case opEQ, opGT, opGTE, opLT, opLTE, opNE:
			if !opCompareFunc(s, l, kop, v, reg) {
				return false
			}
		case opIn:
			if !opInfunc(s, l, kop, v, reg) {
				return false
			}
		case opNin:
			if !opNinfunc(s, l, kop, v, reg) {
				return false
			}
		case opRegex:
			if !opRegexFunc(s, l, kop, v, reg) {
				return false
			}
		case opContain:
			if !opContainFunc(s, l, kop, v, reg) {
				return false
			}

		default:
			panic("$depend right OP can't be " + k)
		}
	}
	return true
}
