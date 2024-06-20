package must

import (
	"errors"
	"testing"
)

type data struct {
	v   any
	err error
}

var testData = data{
	v:   1,
	err: errors.New(""),
}

func TestMust(t *testing.T) {
	tests := []struct {
		name  string
		f     func() (any, error)
		panic bool
	}{
		{
			name: "ok",
			f: func() (any, error) {
				return testData.v, nil
			},
			panic: false,
		},
		{
			name: "panic",
			f: func() (any, error) {
				return nil, testData.err
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); tt.panic != (r != nil) {
					t.Errorf("expected panic: %t, got: %v", tt.panic, r)
				}
			}()

			if v := Must(tt.f()); v != testData.v {
				t.Errorf("expected v: %v, got: %v", testData.v, v)
			}
		})
	}
}
