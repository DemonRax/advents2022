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

func Last(ss []string) int {
	return len(ss) - 1
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
