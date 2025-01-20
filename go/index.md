## Go Fundamentals Notes

### Variables and Simple Data Types

#### Pointers
- Use (value type) values if you want to copy data
```go
a := 42
b := a
a = 27
b // 42
```

- Use (reference type) pointers if you want to share data
```go
a := 42
b := &a // address of a
a = 27 // dereference a
*b // 27
```


```go
a := foo  // create a string
b := &a   // address operator returns address of a variable
*b = "bar" // dereference a pointer with asterick
c = new(int) // built in "new" function creates a pointer to anonymous variable
```

### Data Structures

#### Arrays
- Array and slice indices start at 0.
- static data type 

Declaration 
```go
var arr [3]int // size declared first
fmt.Println(arr) // [0 0 0] zero value default value
arr = [3]int{1, 2, 3} //  [1 2 3]
arr[1] = 99
fmt.Println(arr) // [1 99 3]
fmt.Println(len(arr)) // 3 

arr := [3]string{"foo", "bar", "baz"}
arr2 := arr
fmt.Println(arr2)
arr[0] = "quux"
fmt.Println(arr) // {"quux", "bar", "baz"}
fmt.Println(arr2) // {"foo", "bar", "baz"}
arr == arr2 // false - arrays are comparable

```

#### Slices
- Slice indices start at 0.
- Dynamic data type 

Declaration and manipulation

```go
var s []int // size declared first
fmt.Println(s)
s = []int{1,2,3}
fmt.Println(s[1]) // 2
s[1] = 99
fmt.Println(s) // [1 99 3]

s = append(s, 5, 10, 15) // add elements to the slice
fmt.Println(s) // 1 99 3 5 10 15 

// Delete function requires slices stdlib remove ekements by index specificied
s = slices.Delete(s, 1, 3)
fmt.Println(2) // 1 5 10 15 

s := []string{"foo", "bar", "baz"}
s2 := s  // slices are copied by reference
         // use slices.Clone to copy by value
s[0], s2[2] = "qux", "fred"
fmt.Println(s, s2) //  ["qux", "bar", "fred"] ["qux", "bar", "fred"]
                   // data is shared
s == s2 // compile time error - slices are NOT comparable
        // use Slices
```


#### Maps

```go
var m map[string]int    // declare a map string key with int value
fmt.Println(m)          // nil map
m = map[string]int{"foo":1, "bar":2} // literal map
fmt.Println(m)          // map[foo:1 bar:2]
fmt.Println(m["foo"])   // lookup value in a map
m["bar"] = 99           // update value using same syntax
delete(m, "foo")        // remove entry from map
m["baz"] = 418          // add new value to map
fmt.Println(m)          // map[bar:99 baz:418]
fmt.Println(m["foo])    // returns 0 for unknown key
v, ok := m["foo"]       // comma okay syntax verified presence (T if present F if not)
fmt.Println(v, ok)      // 0, false


m:= map[string]int{
  "foo":1,
  "bar":2,
  "baz":3
}

m2 := m                 // maps are copied by reference
                        // maps.Clone to clone

m["foo"], m2["bar"] = 99, 42
fmt.Println(m)          // map[foo:99 bar:42 baz:3]
fmt.Println(m2)         // map[foo:99 bar:42 baz:3]
                        // data is shared

m == m2                 // compile error - maps are not comparable like slices
```


#### Struct

Fixed size collection but not limited to a single data type.  
1. Field Name
2. Field Type

```go
var s struct {
    name     string
    id       int
}

fmt.Println(s)         // {"" 0} shows us a struct is a value type
s.name = "Arthur"      // assign value to a field
fmt.Println(s.name)    // query value of field

type myStruct struct { // create custom type based on struct
    name     string
    id       int
}
var s myStruct         // declare variable with custom type
fmt.Println(s)         // {"" 0}

s = myStruct{
  name: "Arthur",
  id: 42
}

fmt.Println(s)        // {"Arthur" 42}

s2 := s               // copied by value
s.name = "Tricia"
fmt.Println(s, s2)    // {"Tricia" 42}{"Arthur" 42}
s == s2               // false - since they are value can be compared.
```

### Control Flow: Branches

#### If Statements

```go
if test { ... }


if test { ... 
} else { ...}

if test { ... 
} else if test { ... 
} else { ... }

if initializer; test { ... }


if option == "1" {
  fmt.Println("Option 1")
} else if option == "2" {
  fmt.Println("Option 2")
} else {
  fmt.Println("Option Else")
}
```

#### Case Selects

```go
switch test expression {
  case expression1:
    statements
  case expression2, expression3:
    statements
  default:
    statements
}


i := 5
switch i {
  case 1:
    fmt.Println("First Case")
  case 2 + 3, 2*i+3:
    fmt.Println("Second Case")
  default:
    fmt.Println("Default Case")
}
```


### Control Flow: Loops


#### Loops 

```go
for { ... }                                     // infinite loop
for condition { ... }                           // loop till condition
for initializer; test; post clause { ... }      // counter-based loop

i := 1                                          // infinite loop                              
for {
  fmt.Println(i)
  i += 1
}

i := 1                                          // loop till condition                             
for i < 3 {                                     
  fmt.Println(i)
  i += 1
}
fmt.Println("Done!")

for i:=1;i < 3;i++ {                            // counter-based loop
  fmt.Println(i)
}
```


#### Looping with Collections

- Arrays
    index, value
- Slices
    index, value
- Maps
    key, value

```go
// Options
for key, value := range collection { ... }

for key := range collection { ... }  // Index NOT value

for _, value := range collection { ... }


arr := [3]int{101,102,103}
for i, v := range arr {
        fmt.Println(i, v)
    }
fmt.Println("Done")
```

* Demonstration Program looping with collections

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	type score struct {
		name  string
		score int
	}

	scores := []score{}
	shouldContinue := true

	for shouldContinue {

		fmt.Println("1) Enter a score")
		fmt.Println("2) Print Report")
		fmt.Println("q) Quit")
		fmt.Println("")
		fmt.Println("Please select an option")

		var option string

		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("Enter a student name and score")
			var name, rawScore string
			fmt.Scanln(&name, &rawScore)
			s, _ := strconv.Atoi(rawScore)
			scores = append(scores, score{name: name, score: s})
		case "2":
			fmt.Println("Student Scores")
			fmt.Println(strings.Repeat("-", 14))
			//fmt.Println(scores)
			for _, s := range scores {
				fmt.Println(s.name, s.score)
			}
			fmt.Println("")
		default:
			//exit loop
			shouldContinue = false
		}

	}
}

```

### Functions

- Function Signature


func functionName (parameters)(return values) {

    function body
}

```go
func greet(name string) {
    fmt.Println(name)
}

func greet(name1 string, name2 string) {
    fmt.Println(name1)
    fmt.Println(name2)
}
```

- Variadic Parameters

```go
// receive 0 or more group into a single variable named.. names
// received as a slice
// varies how the caller evokes the function
// Must be final parameter


func greet(names ...string) {
    for _, n := range names {
        fmt.Println(n)
  }
}
```

- Passing Values and Pointers

```go

func main() {
    name, otherName := "Name", "Other name"
    fmt.Println(name)                       // Name
    fmt.Println(otherName)                  // Other name
    myFunc(name, &otherName)                // Does NOT update the pass by value, but does update the pass by reference
    fmt.Println(name)                       // Name
    fmt.Println(otherName)                  // Other new name
}

func myFunc(name string, otherName *string) {
    name = "New name"
    *otherName = "Other new name"
}

```

- Return Values

func functionName (parameters)(**return values**) {
    function body
}

* Returning Single Value

```go

func main() {
  result := add(1, 2)
  fmt.Println(result) // 3

}

func add(l, r int) int {
  return l + r
}
```

* Returning Multiple Values

```go

func main() {
  result, ok := divide(1,2)
  if ok {
    fmt.Println(result)
  }
}

func divide(l, r int) (int, bool) { // return int and bool
  if r == 0 {
    return 0, false
  }
  return l/r, true
}

```

* Named Return Values

```go
func main() {
  result, ok := divide(1,2)
  if ok {
    fmt.Println(result)
  }
}

func divide(l, r int) (result int, ok bool) { 
  if r == 0 {
    return                                    // 0, false
  }
  result = l/r 
  ok = true
  return                                      // optional: return l/r, true
}

```

* Demonstration Program Functions


![[demo-function.go]]

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)
	
type score struct {
  name  string
  score int
}

func main() {

	scores := []score{}
	shouldContinue := true

	for shouldContinue {

		fmt.Println("1) Enter a score")
		fmt.Println("2) Print Report")
		fmt.Println("q) Quit")
		fmt.Println("")
		fmt.Println("Please select an option")

		var option string

		fmt.Scanln(&option)

		switch option {
		case "1":
			scores = append(scores, addStudent())
		case "2":
      printReport(scores) 
		default:
			//exit loop
			shouldContinue = false
		}

	}
}
func addStudent() score {

  fmt.Println("Enter a student name and score")
  var name, rawScore string
  fmt.Scanln(&name, &rawScore)
  s, _ := strconv.Atoi(rawScore)
  return score{name: name, score: s}

}

func printReport(scores []score) {

  fmt.Println("Student Scores")
  fmt.Println(strings.Repeat("-", 14))
  //fmt.Println(scores)
  for _, s := range scores {
    fmt.Println(s.name, s.score)
  }
  fmt.Println("")
} 
```