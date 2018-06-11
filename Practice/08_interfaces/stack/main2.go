package main

import "fmt"

type stack interface {
	peek()
	leng()
}

type stringStack struct {
	stk []string
}

func (s *stringStack) push(n ...interface{}) {
	for _, i := range n {
		s.stk = append(s.stk, i.(string))
	}
}

func (s *stringStack) pop() {
	s.stk = s.stk[0 : len(s.stk)-1]
}

func (s stringStack) leng() {
	fmt.Println(len(s.stk))
}

func (s stringStack) peek() {
	fmt.Println(s.stk[len(s.stk)-1])
}

type intStack struct {
	stk []int
}

func (s *intStack) push(n ...interface{}) {
	for _, i := range n {
		s.stk = append(s.stk, i.(int))
	}
}

func (s *intStack) pop() {
	s.stk = s.stk[0 : len(s.stk)-1]
}

func (s intStack) leng() {
	fmt.Println(len(s.stk))
}

func (s intStack) peek() {
	fmt.Println(s.stk[len(s.stk)-1])
}

func newIntStack() intStack {
	return intStack{}
}
func newStringStack() stringStack {
	return stringStack{}
}

func info(s stack) {
	s.leng()
	s.peek()
}

func main() {
	s := newStringStack()
	i := newIntStack()
	i.push(1, 2, 3, 4, 5, 6)
	s.push("Hello", "World", "Today")
	info(s)
	info(i)

}
