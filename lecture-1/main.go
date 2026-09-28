package main

// package -> module/path.../package_name
import "fmt"

// const c = <value> / expressions

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
// func append(s []T, elem T...) []T -> append(a, 5, 7, 8)
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

// Errors:
// panic ~ exceptions, if not handled lead to crash
// error

/*
type error interface {
	Error() string
}
*/

type Vertex struct {
	x float64
	y float64
}

func NewVertex(x)

func main() {
	var a = [3]int{1, 5, 7}
	var s = a[:]

	a2 := a[2]
	x := 8
	y := 1.2
	c := 1 + 2i

	fmt.Println(s, a2, x, y, c)
}
