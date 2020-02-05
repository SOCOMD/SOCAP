package ocap

import (
	"reflect"
	"testing"
)

func TestNewUnitParser(t *testing.T) {
	tests := []struct {
		Name     string
		Example  string
		Expected rvNewUnit
	}{
		{
			Name:    "Basic new unit",
			Example: `3,6,"Dummy Name 0123","Alpha 1-1","WEST",1`,
			Expected: rvNewUnit{
				Frame:    3,
				ID:       6,
				Name:     "Dummy Name 0123",
				Group:    "Alpha 1-1",
				Side:     "WEST",
				IsPlayer: 1,
			},
		},
	}
	for _, tCase := range tests {
		r, err := rvNewUnitParser(tCase.Example)
		if err != nil {
			t.Fatalf("Failed Test: %s, Err: %s", tCase.Name, err)
		}
		if reflect.DeepEqual(r, tCase.Expected) == false {
			t.Fatalf("Failed Test: %s - E: %+v A: %+v\n", tCase.Name, tCase.Expected, r)
		}

	}
}
