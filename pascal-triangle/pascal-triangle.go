package pascal_triangle

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	result := [][]int{{1}}
	for i := 0; i < numRows-1; i++ {
		lastLine := result[len(result)-1]
		newLine := []int{1}

		for j := 0; j < len(lastLine)-1; j++ {
			newLine = append(newLine, lastLine[j]+lastLine[j+1])
		}
		newLine = append(newLine, 1)
		result = append(result, newLine)
	}

	return result
}
