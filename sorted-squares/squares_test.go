package sorted_squares

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "case 1",
			input: []int{-4, -1, 0, 3, 10},
			want:  []int{0, 1, 9, 16, 100},
		},
		{
			name:  "case 2",
			input: []int{-7, -3, 2, 3, 11},
			want:  []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
			if got := sortedSquaresV2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquaresV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
