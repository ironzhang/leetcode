package stackmac

import (
	"errors"
	"fmt"
)

type Operation string

const (
	Push Operation = "push"
	Ret  Operation = "return"
	Add  Operation = "add"
	Sub  Operation = "sub"
	Mul  Operation = "mul"
	Div  Operation = "div"
)

var (
	ErrStackEmpty = errors.New("stack is empty")
	ErrReturn     = errors.New("return")
)

type Instruction struct {
	OP    Operation
	Value int
}

type Stack struct {
	values []int
}

func (s *Stack) Push(x int) {
	s.values = append(s.values, x)
}

func (s *Stack) Pop() (x int, ok bool) {
	n := len(s.values)
	if n <= 0 {
		return 0, false
	}
	x = s.values[n-1]
	s.values = s.values[:n-1]
	return x, true
}

var optab = map[Operation]func(*Machine, Instruction) (bool, error){
	Push: doPush,
	Ret:  doRet,
	Add:  doAdd,
	Sub:  doSub,
	Mul:  doMul,
	Div:  doDiv,
}

type Machine struct {
	stack Stack
}

func (m *Machine) Run(instructions []Instruction) (int, error) {
	for _, i := range instructions {
		next, err := m.do(i)
		if err != nil {
			return 0, err
		}
		if !next {
			break
		}
	}
	x, ok := m.stack.Pop()
	if !ok {
		return 0, ErrStackEmpty
	}
	return x, nil
}

func (m *Machine) do(i Instruction) (bool, error) {
	op, ok := optab[i.OP]
	if !ok {
		return false, fmt.Errorf("unknown instruction operation %q", i.OP)
	}
	return op(m, i)
}

func doPush(m *Machine, i Instruction) (bool, error) {
	m.stack.Push(i.Value)
	return true, nil
}

func doRet(m *Machine, i Instruction) (bool, error) {
	return false, nil
}

func doAdd(m *Machine, i Instruction) (bool, error) {
	a, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	b, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	c := a + b
	m.stack.Push(c)
	return true, nil
}

func doSub(m *Machine, i Instruction) (bool, error) {
	a, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	b, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	c := b - a
	m.stack.Push(c)
	return true, nil
}

func doMul(m *Machine, i Instruction) (bool, error) {
	a, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	b, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	c := a * b
	m.stack.Push(c)
	return true, nil
}

func doDiv(m *Machine, i Instruction) (bool, error) {
	a, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	b, ok := m.stack.Pop()
	if !ok {
		return false, ErrStackEmpty
	}
	c := b / a
	m.stack.Push(c)
	return true, nil
}
