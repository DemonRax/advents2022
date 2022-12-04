package util

import "strconv"

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Interval []int

func (ii Interval) After(jj Interval) bool {
	if len(ii) < 1 || len(jj) < 1 {
		return false
	}
	return ii[0] > jj[len(jj)-1]
}

func (ii Interval) Before(jj Interval) bool {
	if len(ii) < 1 || len(jj) < 1 {
		return false
	}
	return ii[len(ii)-1] < jj[0]
}
