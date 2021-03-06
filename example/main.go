package main

import (
	"fmt"

	"github.com/pozen/cond"
)

func main() {
	// define a condition:
	// { "key1": "123", "key2": {"$gt": 100} } ,
	c := cond.Cond{
		"key1": "123",
		"key2": cond.Cond{"$gt": 100},
	}

	// create a cond state
	s := cond.NewState().SetCond(c)

	// example1:  expect false
	val_to_check := map[string]interface{}{
		"key1": "123",
		"key2": 99,
	}
	check_result := s.Exec(val_to_check)
	fmt.Printf("check_result is: %v\n", check_result)

	// example2: expect true
	val_to_check = map[string]interface{}{
		"key1": "123",
		"key2": 200,
	}
	check_result = s.Exec(val_to_check)
	fmt.Printf("check_result is: %v\n", check_result)

	// example3: regex
	c2 := cond.Cond{
		"Key1.Key2": cond.Cond{"$regex": "hello|go"},
	}
	// struct as the value to check
	type TVal struct {
		Key1 struct {
			Key2 string
		}
	}
	var val_to_check2 TVal
	val_to_check2.Key1.Key2 = "let's go!"
	// reset cond & check
	check_result = s.SetCond(c2).Exec(&val_to_check2)
	fmt.Printf("regex check_result is: %v\n", check_result)

	// example4: varible
	val_to_check = map[string]interface{}{
		"key1": 123,
		"key2": 200,
	}
	c3 := cond.Cond{
		"key1": cond.Cond{"$lt": "&key2"},
	}
	check_result = s.SetCond(c3).Exec(val_to_check)
	fmt.Printf("variable check_result is: %v\n", check_result)
}
