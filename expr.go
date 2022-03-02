package cond

type ExprOperator struct {
	OP    string
	Prior int
}

var _exprOperators []ExprOperator

func init() {
	_exprOperators = []ExprOperator{
		ExprOperator{OP: "+", Prior: 1},
		ExprOperator{OP: "-", Prior: 1},
		ExprOperator{OP: "*", Prior: 1},
		ExprOperator{OP: "/", Prior: 1},
		ExprOperator{OP: "%", Prior: 1},
		ExprOperator{OP: "^", Prior: 1},
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
