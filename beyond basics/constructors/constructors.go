package main

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f	1
}

type e interface{}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}

func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	mf := Map(m, mult2)
	sf := Map(s, mult2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

func main() {

}

func TestEven(t *testing.T) {
    if Even(2) {
        t.Log("2 should be odd!")
        t.Fail()
    }
}package even 1

import "testing" 2

func TestEven(t *testing.T) { 3
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}func ExampleEven() {
    if Even(2) {
        fmt.Printf("Is even\n")
    }
    // Output: 1
    // Is even
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Stack) pop() (ret int) {
	s.i--
	if s.i < 0 {
		s.i = 0
		return
	}
	ret = s.data[s.i]
	return
}

func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				st.push(r)
				token = ""
			case c == '+':
				fmt.Printf("%d\n", st.pop()+st.pop())
			case c == '*':
				fmt.Printf("%d\n", st.pop()*st.pop())
			case c == '-':
				p := st.pop()
				q := st.pop()
				fmt.Printf("%d\n", q-p)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock impl. */ }
func (m *Mutex) Unlock()  { /* Unlock impl. */ }
func main() {
    k1 := vector.IntVector{}
    k2 := &vector.IntVector{}
    k3 := new(vector.IntVector)
    k1.Push(2)
    k2.Push(3)
    k3.Push(4)
}
type S struct { i int }
func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }
type R struct { i int }
func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }
func f(p I) {
    switch t := p.(type) { 1
        case *S: 2
        case *R: 2
        default: 3
    }
}
func sort(i []interface{}) {  1
    switch i.(type) {         2
    case string:              3
        // ...
    case int:
        // ...
    }
    return /* ... */          4
}func (p Xi) Len() int               {return len(p)}
func (p Xi) Less(i int, j int) bool {return p[j] < p[i]}
func (p Xi) Swap(i int, j int)      {p[i], p[j] = p[j], p[i]}
func (p Xs) Len() int               {return len(p)}
func (p Xs) Less(i int, j int) bool {return p[j] < p[i]}
func (p Xs) Swap(i int, j int)      {p[i], p[j] = p[j], p[i]}
func Sort(x Sorter) { 1
    for i := 0; i < x.Len() - 1; i++ { 2
        for j := i + 1; j < x.Len(); j++ {
            if x.Less(i, j) {
                x.Swap(i, j)
            }
        }
    }
}ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
strings := Xs{"nut", "ape", "elephant", "zoo", "go"}


Sort(ints)
fmt.Printf("%v\n", ints)
Sort(strings)
fmt.Printf("%v\n", strings)type Interface interface {
    sort.Interface
    Push(x interface{})
    Pop() interface{}
}type Person struct {
    name string "namestr"
    age  int
}

func ShowTag(i interface{}) { 1
    switch t := reflect.TypeOf(i); t.Kind() {
    case reflect.Ptr: 2
        tag := t.Elem().Field(0).Tag
	//             <<3>>     <<4>>       <<5>>
	type e interface{}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}

func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	mf := Map(m, mult2)
	sf := Map(s, mult2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

func main() {

}

func TestEven(t *testing.T) {
    if Even(2) {
        t.Log("2 should be odd!")
        t.Fail()
    }
}package even 1

import "testing" 2

func TestEven(t *testing.T) { 3
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}func ExampleEven() {
    if Even(2) {
        fmt.Printf("Is even\n")
    }
    // Output: 1
    // Is even
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Stack) pop() (ret int) {
	s.i--
	if s.i < 0 {
		s.i = 0
		return
	}
	ret = s.data[s.i]
	return
}