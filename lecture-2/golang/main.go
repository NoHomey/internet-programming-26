package main

// package -> module/path.../packagename
import (
	"fmt"
	"math"
)

// const c = <value> / expressions | non exported or private
// const SomeConst = .... | exported or public

// types:
// int, uint (arch dependent): 0
// int8, int16, int32, int64: 0
// uint8, uint16, uint32, uint64: 0
// float32, float64: 0.0
// complex64, complex128 (32bit, 64bit)
// str (string, utf-8, run arrays, immutable): ""
// rune = unt32: 0
// bool: false

// [n]<Type> - static array of type <Type> with length n.
// []<Type> - dynamic array (or slice) of type <Type> (they have length and capacity)
// a = [1, 5, 7]
// a[0:3] -> [1, 5, 7]
// a[1:3] -> [5, 7]
// a[0:2] -> [1, 5]
// a[0:] -> [1, 5, 7]
// a[:3] -> [1, 5, 7]
// a[:] -> [1, 5, 7]
// a[i] -> elem at index i
// a[i:i+1] -> [a[i]]
// make([]T, N) -> allocates slice with length and capacity of N and type T
// make([]T, 0, N) -> empty slice (with length 0) and capacity N and type T
// make([]int, 0, 5) -> [0, 0, 0, 0, 0] (length=0, capacity=5)
// make([]int, 3, 5) -> [0, 0, 0, 0, 0] (length=3, capacity=5)
// func append(s []T, elem ...T) []T -> append(a, 5, 7, 8)
// len(s []T) int -> returns the length
// cap(s []T) int -> returns the capacity
// copy(x []T, y []T)

// [3][3]float64
// [][]int -> table
// map[str, []float32]
// map[*T, []int]

// map[K, V], K is comparable, ==, != (HashMap)
// var m map[str, int] (zero value is nil)
// m["abc"] = 9 // add or assign/change
// delete(m, "abc") // deletes the entry for "abc"
// v, ok := m["abc"] // if ok is true then there is element on key "abc" and v will be the value of this element
//                   // if ok is false then there is no element at key "abc" and v will have zero value (in this case 0)
// x := m["abc"] // get
// len(m) -> current size
// cap(m) -> capacity of the map

// set ~ map[Elem, bool]

// var x int8 = 5
// int16(x) -> produces new value, explicit type casting | T(v)

// *T (pointer of type T)
// &v -> address of v
// p := &v
// *p -> direct access to the memory (value) pointed by p

/* type Slice[T comparable] struct {
	length int
	capacity int
	mem *T
}
*/

/*
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	--- while loop
	i := 0
	for i < n {
		fmt.Println(i)
		i++
	}

	// infinite loop or loop with dynamic condition or multiple exit conditions

	for {
	 ...

	 if condition {
	 	break
	 }
	}

	for x := exp, y := something else; condition(x, y) {

	} else {
	 // here x and y are defined and accessible
	}
*/

// Errors:
// panic ~ exceptions, if not handled lead to crash
// error

/*
type error interface {
	Error() string
}
*/

/*

type Stringer interface {
	String() string
}
*/

type Vertex struct {
	x float64
	y float64
}

func NewVertex(x, y float64) Vertex {
	return Vertex{x: x, y: y}
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(c float64) {
	v.x *= c // (*v).x *= c
	v.y *= c // (*v).y *= c
}

func (v Vertex) String() string {
	return fmt.Sprintf("(x=%f, y=%f)", v.x, v.y)
}

func myPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("%d\n", x.(int))
	case float64:
		// fmt.Printf("%d", x.(int)) this will panic because x is not int here
		fmt.Printf("%f\n", x.(float64))
	case string:
		fmt.Println(x.(string))
	case fmt.Stringer:
		myPrint(x.(fmt.Stringer).String())
	case error:
		myPrint(x.(error).Error())
	}
}

// Type aliasing
type myInt = int16
type myFunc = func(int) int

// New Type
type mySpecialVertex Vertex

func (v mySpecialVertex) Abs() float64 {
	return math.Abs(v.x) + math.Abs(v.y)
}

func main() {
	v := NewVertex(1, 1)
	v.Scale(4)
	fmt.Println(v)
	fmt.Println(v.Abs())

	x := 9

	switch x {
	case 2:
		fmt.Println("x is 2")
		fallthrough
	case 9:
		fmt.Println("Our value")
	case 3, 5:
		fmt.Println("Something")

	default:
		fmt.Println("Unknown")
	}

	vabs := v.Abs

	fmt.Println(vabs())

	myPrint(x)
	myPrint(v)

	var s any
	s = v
	if vec, ok := s.(Vertex); ok {
		fmt.Println(vec.x)
	}

	var m mySpecialVertex = mySpecialVertex(v)
	fmt.Println(m.Abs(), Vertex(m).Abs())
}
