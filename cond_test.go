package cond

import (
	"testing"
)

func TestCond1(t *testing.T) {
	var t1, t2 struct {
		LV1 struct {
			LV2_1 string
			LV2_2 struct {
				LV3_1 string
				LV3_2 bool
			}
		}
	}

	t1.LV1.LV2_1 = "123"
	t2.LV1.LV2_2.LV3_1 = "12"
	t2.LV1.LV2_2.LV3_2 = true

	tests := []struct {
		dsc    string
		c      Cond
		val    interface{}
		expect bool
	}{
		{"regex match", Cond{"Key": Cond{"$regex": "Go|Hello"}}, struct{ Key string }{Key: "123 Go 34"}, true},
		{"regex dismatch", Cond{"Key": Cond{"$regex": "Hello|Go"}}, struct{ Key string }{Key: "let's g!"}, false},
		{"eq num", Cond{"Key": 123}, struct{ Key int }{Key: 123}, true},
		{"eq num 2", Cond{"Key": 123.456}, struct{ Key float32 }{Key: 123.456}, true},
		{"eq num 3", Cond{"Key": 123}, struct{ Key uint16 }{Key: 123}, true},

		{"ne num", Cond{"Key": 1234}, struct{ Key int }{Key: 123}, false},
		{"ne num 2", Cond{"Key": 123.456}, struct{ Key float64 }{Key: 123.4567}, false},
		{"ne num 3", Cond{"Key": 123}, struct{ Key uint16 }{Key: 12}, false},
		{"eq string", map[string]interface{}{"Key": "123"}, struct{ Key string }{Key: "123"}, true},
		{"ne string", map[string]interface{}{"Key": "123"}, struct{ Key string }{Key: "12"}, false},

		{"eq bool", map[string]interface{}{"Key": true}, struct{ Key bool }{Key: true}, true},
		{"ne bool", map[string]interface{}{"Key": true}, struct{ Key bool }{Key: false}, false},

		{"eq nest", map[string]interface{}{"LV1.LV2_1": "123"}, &t1, true},
		{"ne nest true", map[string]interface{}{"LV1.LV2_2.LV3_1": "12"}, &t2, true},
		{"ne nest false", map[string]interface{}{"LV1.LV2_2.LV3_1": "123"}, &t2, false},
		{"bool eq nest", map[string]interface{}{"LV1.LV2_2.LV3_2": true}, &t2, true},
		{"bool ne nest", map[string]interface{}{"LV1.LV2_2.LV3_2": false}, &t2, false},
		{"bool ne nest2", map[string]interface{}{"LV1.LV2_2.LV3_2": false}, &t1, true},
		{"bool ne nest3", map[string]interface{}{"LV1.LV2_2.LV3_2": true}, &t1, false},
		{"eq struct", map[string]interface{}{"Key": true, "Key2": 123}, struct {
			Key  bool
			Key2 int
		}{Key: true, Key2: 123}, true},

		{"ne struct", map[string]interface{}{"Key": true, "Key2": 123}, struct {
			Key  bool
			Key2 int
		}{Key: true, Key2: 12}, false},

		{"and", map[string]interface{}{
			"$and": []map[string]interface{}{
				map[string]interface{}{"Key": true},
				map[string]interface{}{"Key2": 123},
			},
		}, struct {
			Key  bool
			Key2 int
		}{Key: true, Key2: 123}, true},

		{"and 2", map[string]interface{}{
			"$and": []map[string]interface{}{
				map[string]interface{}{"Key": true},
				map[string]interface{}{"Key2": 123},
			},
		}, struct {
			Key  bool
			Key2 int
		}{Key: true, Key2: 1234}, false},

		{"or", map[string]interface{}{
			"$or": []map[string]interface{}{
				map[string]interface{}{"Key": true},
				map[string]interface{}{"Key2": 123},
			},
		}, struct {
			Key  bool
			Key2 int
		}{Key: true}, true},

		{"or 2", map[string]interface{}{
			"$or": []map[string]interface{}{
				map[string]interface{}{"Key": true},
				map[string]interface{}{"Key2": 123},
			},
		}, struct {
			Key  bool
			Key2 int
		}{Key: false}, false},

		{"gt", map[string]interface{}{
			"Key": map[string]interface{}{"$gt": 100},
		}, struct {
			Key int
		}{Key: 1234}, true},

		{"gt 2", map[string]interface{}{
			"Key": map[string]interface{}{"$gt": 10000},
		}, struct {
			Key int
		}{Key: 1234}, false},

		{"in 2", map[string]interface{}{
			"Key": map[string]interface{}{"$in": []string{"1234", "1", "3"}},
		}, struct {
			Key string
		}{Key: "1"}, true},

		{"in 3", map[string]interface{}{
			"Key": map[string]interface{}{"$in": []string{"1234", "1", "3"}},
		}, struct {
			Key string
		}{Key: "2"}, false},
		{"contain 2", map[string]interface{}{
			"Key": map[string]interface{}{"$contain": "2"},
		}, struct {
			Key []string
		}{Key: []string{"2", "1", "3"}}, true},
		{"contain 4", map[string]interface{}{
			"Key": map[string]interface{}{"$contain": 4},
		}, struct {
			Key []int
		}{Key: []int{2, 1, 4}}, true},

		{"not contain 4", map[string]interface{}{
			"Key": map[string]interface{}{"$contain": 4},
		}, struct {
			Key []int
		}{Key: []int{2, 1}}, false},

		{
			dsc: "and or nest",
			c: map[string]interface{}{
				"$and": []map[string]interface{}{
					{
						"Key": map[string]interface{}{
							"$gt": 0,
						},
					},
					{
						"$or": []map[string]interface{}{
							{
								"Key": 0,
							},
							{
								"Key": 1,
							},
						},
					},
				},
			},
			val: struct {
				Key int
			}{Key: 1},
			expect: true,
		},

		{
			dsc: " or nest",
			c: map[string]interface{}{
				"Key": 1,
				"$or": []map[string]interface{}{
					{
						"Key": map[string]interface{}{
							"$lt": 0,
						},
					},
					{
						"Key": map[string]interface{}{
							"$gt": 0,
						},
					},
				},
			},
			val: struct {
				Key int
			}{Key: 1},
			expect: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dsc, func(t *testing.T) {
			s := NewState().SetCond(tt.c)
			if got := s.Exec(tt.val); got != tt.expect {
				t.Errorf("Cond.Exec() = %v, expect = %v, %+v", got, tt.expect, tt)
			}
		})
	}
}
