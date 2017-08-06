package pascal

const testVersion = 1

// Triangle generates pascal triangle with `n` rows.
func Triangle(n int) (res [][]int) {
	res = make([][]int, n)

	for rowIdx := range res {
		row := make([]int, rowIdx+1)
		res[rowIdx] = row

		row[len(row)-1] = 1
		row[0] = 1
		for i := 1; i < rowIdx; i++ {
			row[i] = res[rowIdx-1][i-1] + res[rowIdx-1][i]
		}
	}

	return
}
