package cond

type ExprOperator struct {
	OP    string
	Prior int
	L     bool
	R     bool
}

var _exprOperators []ExprOperator

func init() {
	_exprOperators = []ExprOperator{
		ExprOperator{OP: "+", Prior: 1, L: true, R: true},
		ExprOperator{OP: "-", Prior: 1, L: true, R: true},
		ExprOperator{OP: "*", Prior: 1, L: true, R: true},
		ExprOperator{OP: "/", Prior: 1, L: true, R: true},
		ExprOperator{OP: "%", Prior: 1, L: true, R: true},
		ExprOperator{OP: "^", Prior: 1, L: true, R: true},
		ExprOperator{OP: "$", Prior: 1},
	}
}

func matchExprOperator(c string) *ExprOperator {
	for i, v := range _exprOperators {
		if v.OP == c {
			return &_exprOperators[i]
		}
	}
	return nil
}
