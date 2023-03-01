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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMountainRemovals(tt.nums); got != tt.want {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
