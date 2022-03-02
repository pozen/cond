package cond

import (
	"regexp"
)

const opRegex Operator = "$regex"

func opRegexFunc(s *State, l string, op Operator, r1 interface{}, reg map[string]interface{}) bool {

	r := s.loadVarible(r1, reg)

	regexStr, ok := r.(string)
	if !ok {
		panic("$regex right value type must be string")
	}

	lval := s.loadValFromSession(l, reg)
	if lval == nil {
		return false
	}

	srcStr, ok2 := lval.(string)
	if !ok2 {
		return false
	}

	regexc, err := regexp.Compile(regexStr)
	if err != nil {
		panic("wrong regex synctax : " + regexStr)
	}

	return regexc.MatchString(srcStr)
}
