package leetcode

import (
	"reflect"
	"testing"
)

func TestGetFolderNames(t *testing.T) {
	tests := []struct {
		inputs  []string
		outputs []string
	}{
		{
			inputs:  []string{"pes", "fifa", "gta", "pes(2019)"},
			outputs: []string{"pes", "fifa", "gta", "pes(2019)"},
		},
		{
			inputs:  []string{"gta", "gta(1)", "gta", "avalon"},
			outputs: []string{"gta", "gta(1)", "gta(2)", "avalon"},
		},
		{
			inputs:  []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"},
			outputs: []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece(4)"},
		},
		{
			inputs:  []string{"wano", "wano", "wano", "wano"},
			outputs: []string{"wano", "wano(1)", "wano(2)", "wano(3)"},
		},
		{
			inputs:  []string{"kaido", "kaido(1)", "kaido", "kaido(1)"},
			outputs: []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"},
		},
	}
	for _, tt := range tests {
		if got, want := getFolderNames(tt.inputs), tt.outputs; !reflect.DeepEqual(got, want) {
			t.Fatalf("get folder names: got %v, want %v", got, want)
		}
	}
}
