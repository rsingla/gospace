package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_validate_random(t *testing.T) {
	type args struct {
		Value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate_random(tt.args.Value); got != tt.want {
				t.Errorf("validate_random() = %v, want %v", got, tt.want)
			}
		})
	}
}
