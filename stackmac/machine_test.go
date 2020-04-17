package stackmac

import "testing"

func TestMachine(t *testing.T) {
	instructions := []Instruction{
		{OP: Push, Value: 10},
		{OP: Push, Value: 20},
		{OP: Add},
		{OP: Push, Value: 40},
		{OP: Sub},
		{OP: Push, Value: 3},
		{OP: Mul},
		{OP: Push, Value: -2},
		{OP: Div},
	}
	expect := 15

	var m Machine
	ret, err := m.Run(instructions)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if ret != expect {
		t.Fatalf("return %d, want %d", ret, expect)
	}
	t.Logf("return %d", ret)
}
