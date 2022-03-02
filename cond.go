package cond

import (
	"context"
	"strings"

	"github.com/fatih/structs"
)

// support simple condtion query syntax , just like mongo json query.
// eg:
//	{
//		"a": "123",
//		"a.b": {"$gt":0, "$lt":10},
//		"$or": [{"a":1},{"c":{"$ne": "123"}}]
//	}

// condtion define
type Cond map[string]interface{}

type State struct {
	Ctx                 context.Context
	Cond                Cond
	NestedKeySplitMark  string
	RuntimeVariableMark string
}

func NewState() *State {
	return &State{
		Ctx:                 context.Background(),
		NestedKeySplitMark:  ".",
		RuntimeVariableMark: "&",
	}
}

func (s *State) SetCond(c Cond) *State {
	s.Cond = c
	return s
}

func (s *State) SetNestedKeySplitMark(m string) *State {
	s.NestedKeySplitMark = m
	return s
}

func (s *State) Exec(session interface{}) bool {
	var reg map[string]interface{}
	var ok bool
	if reg, ok = session.(map[string]interface{}); !ok {
		reg = structs.Map(session)
	}
	return opMainFunc(s, "", opMain, s.Cond, reg)
}

// load key (condtion's) value from runtime session
// if state.nestedKeySplitMark is ".", then the key of condition should be like this :
// cond = {
//		"level1.level2.level3": "some value"
// },
// and the key "level1.level2.level3" will try to match the session value like this:
// session =	{
//		"level1": {
//			"level2": {
//				"level3" : "it is me!"
//			}
//		}
//	}
//
// s.loadValFromSession("level1.level2.level3", session) returns "it is me!"
//
func (s *State) loadValFromSession(k string, reg map[string]interface{}) interface{} {
	keys := strings.Split(k, s.NestedKeySplitMark)
	tmp := reg
	for i, key := range keys {
		v, ok := tmp[key]
		if !ok {
			return nil
		}
		if i >= len(keys)-1 {
			return v
		}
		if tmp, ok = v.(map[string]interface{}); !ok {
			return nil
		}
	}
	return nil
}

func (s *State) loadVarible(r interface{}, reg map[string]interface{}) interface{} {
	str, ok := r.(string)
	if !ok {
		return r
	}
	if strings.HasPrefix(str, s.RuntimeVariableMark) {
		str = strings.TrimPrefix(str, s.RuntimeVariableMark)
	} else {
		return r
	}
	return s.loadValFromSession(str, reg)
}
