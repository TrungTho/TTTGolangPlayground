package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
		// assert(equal, 1, 1)
	}
}

func Test_hehe(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "ok",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hehe(); got != tt.want {
				t.Errorf("hehe() = %v, want %v", got, tt.want)
			}
		})
	}
}
