package problem

func permutations(list []int) [][]int {
	size := len(list)
	if size == 0 {
		return [][]int{[]int{}}
	}

	perms := [][]int{}

	for i := 0; i < size; i++ {
		// The sublist except for element i
		sublist := make([]int, i)
		copy(sublist, list[0:i])
		sublist = append(sublist, list[i+1:]...)

		subperms := permutations(sublist)

		for _, subperm := range subperms {
			perms = append(
				perms,
				append([]int{list[i]}, subperm...),
			)
		}
	}

	return perms
}

func isSolution(a, b, c, d, e, f int) bool {
	return a+c+e+d == 14 && d+b+c+f == 14 && a+b+e+f == 14
}

func SolveProblem001() [][]int {
	perms := permutations([]int{1, 2, 3, 4, 5, 6})

	solutions := [][]int{}

	for _, perm := range perms {
		if isSolution(perm[0], perm[1], perm[2], perm[3], perm[4], perm[5]) {
			solutions = append(solutions, perm)
		}
	}

	return solutions
}
