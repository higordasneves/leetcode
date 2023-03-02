package mountain

import "testing"

func Test_minimumMountainRemovals(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 3, 1},
			want: 0,
		},
		{
			name: "example 2",
			nums: []int{2, 1, 1, 5, 6, 2, 3, 1},
			want: 3,
		},
		{
			name: "example 3",
			nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want: 1,
		},
		{
			name: "example 4",
			nums: []int{4, 3, 2, 1, 1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example 5",
			nums: []int{1, 16, 84, 9, 29, 71, 86, 79, 72, 12},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMountainRemovals(tt.nums); got != tt.want {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
