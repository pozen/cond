**What's this?**   
A lightweight configurable condition rule check lib. Support simple condtion query syntax , just like mongo json query.
eg:

```
{
	"a": "123",
	"a.b": {"$gt":0, "$lt":10},
	"$or": [{"a":1},{"c":{"$ne": "123"}}],
	"c": {"$in": ["1", "2", "3"]},
	"$and": [{"a.c": 100.001}, {"d": {"$regex": "go|home"}}]
}

```

(You can find more examples in unit test file: cond_test.go)

You can use it as a rule engine for condition check (run conditon check over the runtime data). Normally the Cond is a config data from config file or DB or API query.    

Supports operators now:  
|operators    |dsc          |
| ----------- | ----------- | 
|$eq          | eg: {"a": 123}, {"a": {"$eq": 123}} |
|$ne          | {"a": {"$ne": 123}} |
|$gt          | {"a": {"$gt": 123}} |
|$gte          | {"a": {"$gte": 123}} |
|$lt          | {"a": {"$lt": 123}} |
|$lte          | {"a": {"$lte": 123}} |
|$and          | {"$and": [{"a": 123}, {"b": 125}]} |
|$or          | {"$or": [{"a": 123}, {"b": 125}]} |
|$in          | {"a": {"$in": [1,2,3,4,5]} |
|$nin          | {"b": {"$nin": [1,2,3,4,5]} |
|$regex          | {"a": "$regex": "hello|go"} |
|$contain          | {"a": {"$contain": "abc"}} |

**How to use?**  

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

```
