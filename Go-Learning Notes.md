# Go

**Note**: These notes are basically copied from [A Tour of Go](https://go.dev/tour/list).

  

### Packages

Go programs are made up of **packages**

Programs start running in package `main`.

```go
package main

func main() {
    ...
}
```



### Imports

Import statements should be grouped into parenthesis.

```go
import (
    "fmt"	// implements formatted I/Os
    "math"
)
```



### Exported Names

Name is *exported* if it begins with a capital letter.

Example:

```go
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
}
```





### Variables

Adding types to function parameters:

```go
func add(x int, y int) int {
    return x + y
}

// OR

func add(x, y int) int {
    return x + y
}
```





A function can return any number of results

```go
func swap(x, y string) (string, string) {
    return y, x // return two strings
}
```





**Named return values**

Return values may be names, and they will be treated as variables defined at the top of the function.

A `return` statement without arguments returns the named return values. This is known as a **naked return**.

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return	// naked return of (x, y int)
}
```



The `var` statement declares a list of variables, the type is at last.

```go
package main

import "fmt"

var c, python, java bool	// declare 3 boolean variables

func main() {
    var i int
    fmt.Println(i, c, python, java)
}

// output: 0 false false false : the default zero values
```



Variable declaration can include initializers, the type can be omitted.

```go
var i, j int = 1, 2
func main() {
    var c, python, jave = true, false,  "no!"
    fmt.Println(i, j, c, python, java)
}
```



Outside a function, every statement should begin with a keyword (`var`, `func`, etc.) => the `:=` operator is not available.

```go
syntax error: non-declaration statement outside function body
```



Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with *implicit* type.

```go
var n = 4	// must use `var`
func main() {
    var i, j int = 1, 2 // can do this: declaration statement with `var`
    k := 3				// can do this: short assignment statement
    c, python, java := true, false, "no!"
}
```



Variables declared without explicit initial value are given their **zero value**.

```go
numeric types: 0
boolean: false
string: ""



func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// output: 0 0 false ""
```





Type conversions: `T(v)` converts variable `v` to the type `T`.

```go
var i int = 42
var f float64 = float64(i)
// OR
f := float64(i)
```



### Constants

- declared like variables, with the `const` keyword
- cannot be declared using the `:=` syntax

```go
const Truth = true
```

  



***

  



### For Loop

Go has only one looping construct: the `for` loop

```go
func main() {
    sum := 0 
    for i := 0; i < 10; i++ {		// no parentheses
        sum += i
    }
}
```



The init and post statements are optional

```go
func main() {
    sum := 0
    for ; sum < 1000; {
        sum += 1
    }
    fmt.Println(sum)
}
```



For loop for while loop

```go
// just drop the semi colons
func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
```



**Forever**

Drop everything (except `for`) to create the infinite loop

```go
func main() {
    for {
    }
}
```

  





### If Statements

Expressions for if statements need not be surrounded by parentheses (but is can tho)

```go
func main() {  
    if true {
        ...
    }
    
    if x < 10 {
        ...
    }
}
```



If can start with a short statement

```go
func pow(x, n, lim float64) float64 {
    if v:= math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n". v, lim)
    }
    return lim
}
```





### Switch Statement

- runs the first case whose value is equal to the condition expression
- Go only runs the selected case,
- Go will automatically provide the `break` statement at the end of each case (no need to explicitly type out `break`)

```go
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```





### Defer Statement

A defer statement defers the execution of a function until the surrounding function returns.

- deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// output:
hello
world
```

Deferred function calls are pushed onto a slack, executed in last-in-first-out order:

```go
func main() {
    fmt.Println("counting")
    
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("done")
}

// output: from 9 to 0
```





### Pointers

A pointer holds the memory address of a value.

`*T` is a pointer to a `T` value. It's zero value is `nil`.



The `&` operator generates a pointer to its operand.

```go
i := 42
p = &i
```



The `*` operator denotes the pointer's underlying value. (**dereference**)

```go
fmt.Println(*p) // read i through pointer p
*p = 21			// set i through the pointer p
```

Go has no pointer arithmetic.



#### Structs

A `struct` is a collection of fields

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
    p.X = 1e9			// actually is (*p).X
	fmt.Println(v)
}
```





### Arrays

`[n]T` is an array of `n` values of type `T`.

Array's size is fixed.

```go
func main() {}
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// output:
[Hello World]
[2 3 5 7 11 13]
```





### Slices

Slice is a dynamically-sized, flexible view into the elements of an array.

`[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```go
a[low : high]
// high will not be included
// Example:

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}	// declare an array
    var s[]int = primes[1:4]	// declare a slice
    fmt.Println(s)
}
```



A slice does not store any data, it just describe a section of an underlying array.

Slices are like references to arrays.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

```go
func main() {
    names := [4]string{"John", "Paul", "George", "Ringo"}
    fmt.Println(names)
    // output: [John Paul George Ringo]
    
    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)
    // output: [John Paul] [Paul George]
    
    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
    // output: [John XXX] [XXX George]
	//		   [John XXX George Ringo]
}
```



Slice Expressions:

```go
var a [10]int
// These expressions are equivalent
a[0:10]
a[:10]
a[0:]
a[:]
```





Slice has both a **length** and a **capacity**.

Length is the number of elements it contains. (`len(s)`)

Capacity is the number of elements in the underlying array, counting from the first element in the slice. (`cap(s)`)

You can extend s slice's length by re-slicign it, if it has sufficient capacity. (`runtime error: slice bounds out of range`)

```go
s := []int{2, 3, 5, 7, 11, 13}
// len = 6, cap = 6 [2 3 5 7 11 13]

// Slice the slice to give it zero length
s = s[:0]
// len = 0, cap = 6 []

// Extend its length
s = s[:4]
// len = 4, cap = 6 [2 3 5 7]

// Drop its first two values
s = s[2:]
// len = 2, cap = 4 [5 7]

// Extend it again
s = s[:4]
// len = 2, cap = 4 [5 7 11 13]
```



#### Nil Slice

The zero value of a slice is `nil`. A nil slice has a length and capacity of 0, and no underlying array.

```go
var s []int
fmt.Println(s, len(s), cap(s))
// output: [] 0 0
// this is a nil slice
```



#### Creating a slice

Slices can be created with the built-in `make` function. This is how you create dynamically-sized arrays.

The `make` function allocates a zeroed array and returns a slice that refers to that array:

```go
a := make([]int, 5)	// len(a) = 5
// cap(a) = 5
```

Pass a thrid argument to specify a capacity

```go
b := make([]int, 0, 5)	// len(b) = 0, cap(b) = 5
```





Slices can contain any type, including other slices

```go
// create a tic-tac-toe board => 2D slices
board := [][]string{
    []string{"O", "X", "_"},
	[]string{"O", "X", "_"},
	[]string{"_", "_", "X"},
}
```



You can append values to a slice, and it will increase the slide's capacity.

```go
var s[]int
// len = 0, cap = 0 []

s = append(s, 0)
// len = 1, cap = 1 [0]

s = append(s, 1, 2, 3)
// len = 4. cap = 4 [0 1 2 3]
```





#### Range

Range form of `for` loop iterates over a slice.

In each iteration, two values are returned, the first is the *index*, and the second is a copy of that element in the index.

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

You can skip the index or value by assigning to `_`.

```go
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```go
for i := range pow
```





  



### Maps

A map maps keys to values

The zero value of a map is `nil`. A `nil` map has no keys.

The `make` function returns a map of the given type, initialized and ready for use.

```go
type Vertex struct  {
    Lat, Long float64
}

var m map[string]Vertex

m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
}
```



#### Mutating Maps

Insert or update an element:

```go
m[key] = elem
```



Retrieve an element:

```go
elem = m[key]
```



Delete an element:

```go
delete(m, key)
```



Test that a key is present with a two-value assignment:

```go
elem, ok = m[key]
```

- If `key` is in `m`, then `ok` is `true`, otherwise `ok` is `false`.
- if `key` is not in `m`, then `elem` is the zero value for the map's element type

- If `elem` or `ok` have not yet been declared, use a short declaration form:

  ```go
  elem, ok := m[key]
  ```

  

### Function Values

Functions are values, they can be passed around just like other values. They may be used as function arguments and return values.

```go
func compute(fn func(float64, float64) floaat64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(compute(hypot))
    fmt.Println(compute(hypot))
}
```

  

### Function Closures

A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables, in this sense the function is "bound" to the variables.



Example: `adder` function returns a closure, each closure is bound to its own `sum` variable.

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 5; i++ {
        fmt.Println(
            pos(i),
            neg(-2 * i),
        )
    }
}

// output:
0 0
1 -2
3 -6
6 -12
10 -20
```



Implementation of fibonacci numbers (0, 1, 1, 2, 3, 5)

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first + second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// output: 0 1 1 2 3 5 8 13 21 34
```





  



***

  



### Methods

Go does not have classes, but you can define methods of types.

A method is a function with special *receiver* argument. The receiver appears in its own argument list between the `func` keyword and the method name.

```go
type Vertex struct {
    X, Y float64
}

// Abs has a receiver of type Vertex named v
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```



You can declare methods with pointer receivers. Methods with pointer receivers can modify the value to which the receiver points. Pointer receivers are more common than value receivers.

```go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```



We can rewrite `Scale` as function:

```go
func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)	// must pass in a pointer
	fmt.Println(Abs(v))
}
```

- functions with a pointer argumetn must take a pointer

  ```go
  var v Vertex
  Scale(v, 5)		// Compile error!
  Scale(&v, 5)	// OK
  ```

  

- methods with pointer receivers take either a value or a pointer as the receiver when they are called

  ```go
  var v Vertex
  v.Scale(5)		// OK, Go interprets v.Scale(5) as (&v).Scale(5) since the Scale method hs a pointer receiver
  p := &v
  p.Scale(10)		// OK
  ```



Reasons to use pointer receiver:

- the method can modify the value that its receiver points to
- avoid copying the value on each method call

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.



## Interfaces

An *interface type* is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.

```go
var a Abser
f := MyFloat(-math.Sqrt2)
v := Vertex{3, 4}

a = f // a MyFloat implements Abser
a = &v // a *Vertex implements Abser

a = v // a does NOT implement Abser because Abs method is only defined on *Vertex

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```



A type implements the interface by implementing its methods. 

```go
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```



The **empty interface** specifies zero methods. They may hold values of any type.

```go
interface{}
```

Empty interfaces are used by code that handles values of unknown type. (e.g. `fmt.Println`)



#### Type assertions

A **type assertion** asserts that the interface value `i` holds the concerte type `T` and assigns the underlying `T` value to the variable `t`. If `i` does not hold a `T`, the statement will triggers a panic.

```go
t := i.(T)
```



A type assertion can be used to test whether an interface  value hold a specific type. Type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```go
t, ok := i.(T)
```

- if `i` holds a `T`, `t` will be the underlying value and `ok` will be true,
- otherwise, `t` will be the zero value of type `T` and `ok` will be false, no panic occurs.



```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)		// hello

s, ok := i.(string)
fmt.Println(s, ok)	// hello true

f, ok := i.(float64)
fmt.Println(f, ok)	// 0 false

f = i.(float64) 	// panic
fmt.Println(f)
```



  

A **type switch** is a construct that  permits several types assertions in series. Values are compared against the type of the value held by the given interface value. 

```go
switch v := i.(type) {	// keyword type is used here
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```



#### Stringers

A `Stringer` is a type that can describe itself as a string. The `fmt` package look for this interface to print values.

```go
type Stringer interface {
    String() string
}
```



Example: implement `fmt.Stringer` to print the ip address as a dotted quad

```go
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
    // IPAddr{1, 2, 3, 4} should print as "1.2.3.4"
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```



### Error

Go program express error state with `error` values.

The `error` type is a built-in interface

```go
type error interface {
    Error() string
}
// fmt package looks for the error interface when printing values
```

A nill `error` denotes a success, a non-niil `error` denotes failure.



