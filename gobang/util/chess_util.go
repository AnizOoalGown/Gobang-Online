package util

import "gobang/entity"

func HasStep(i int8, j int8, color int8, steps *[]entity.Chess) bool {
	for k := int(color); k < len(*steps); k += 2 {
		step := (*steps)[k]
		if step.I == i && step.J == j {
			return true
		}
	}
	return false
}

func checkFiveInDirection(i int8, j int8, color int8, x int8, y int8, steps *[]entity.Chess) bool {
	count := 1
	for m, n := i-x, j-y; m >= 0 && n >= 0 && m < 15 && n < 15; m, n = m-x, n-y {
		if HasStep(m, n, color, steps) {
			count++
		} else {
			break
		}
	}
	for m, n := i+x, j+y; m >= 0 && n >= 0 && m < 15 && n < 15; m, n = m+x, n+y {
		if HasStep(m, n, color, steps) {
			count++
		} else {
			break
		}
	}

	return count >= 5
}

func CheckFiveOfLastStep(steps *[]entity.Chess) (bool, int8) {
	color := int8((len(*steps) - 1) % 2)
	if len(*steps) < 9 {
		return false, color
	}
	lastStep := (*steps)[len(*steps)-1]
	i := lastStep.I
	j := lastStep.J

	hasFive := checkFiveInDirection(i, j, color, 1, 0,
		steps) || checkFiveInDirection(i, j, color, 0, 1,
		steps) || checkFiveInDirection(i, j, color, 1, 1,
		steps) || checkFiveInDirection(i, j, color, 1, -1, steps)

	return hasFive, color
}
