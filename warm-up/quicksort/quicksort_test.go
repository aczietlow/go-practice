package main

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			name: "empty slice",
			list: []int{},
			want: []int{},
		},
		{
			name: "single element",
			list: []int{1},
			want: []int{1},
		},
		{
			name: "already sorted",
			list: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			list: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "random order",
			list: []int{3, 6, 1, 8, 4, 5},
			want: []int{1, 3, 4, 5, 6, 8},
		},
		{
			name: "with duplicates",
			list: []int{3, 1, 2, 3, 1},
			want: []int{1, 1, 2, 3, 3},
		},
		{
			name: "negative numbers",
			list: []int{-1, -3, 2, 0, 1},
			want: []int{-3, -1, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quicksort(tt.list)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}
