package stock_spanner

type StockSpanner struct {
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for l := len(this.stack) - 1; l > -1 && this.stack[l] <= price; l-- {
		count++
	}
	this.stack = append(this.stack, price)

	return count
}

// ex 901: Online Stock Span v0: https://leetcode.com/problems/online-stock-span/description/
