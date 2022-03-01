***** what's this? **
  
A lightweight configurable condition rule check lib. Support simple condtion query syntax , just like mongo json query.
eg:

```
{
	"a": "123",
	"a.b": {"$gt":0, "$lt":10},
	"$or": [{"a":1},{"c":{"$ne": "123"}}]
}

```

You can use it as a rule engine for condition check, normally the Cond is a config data from configure file or DB.  

** How to use? **  
  
```
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
}

```
