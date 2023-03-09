package stock_spanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockSpanner_Next(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   []int
	}{
		{
			name:   "example 1",
			prices: []int{100, 80, 60, 70, 60, 75, 85},
			want:   []int{1, 1, 1, 2, 1, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for i, price := range tt.prices {
				assert.Equal(t, tt.want[i], this.Next(price))
			}
		})
	}
}
