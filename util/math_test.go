package util

import "testing"

func Test_Intervals(t *testing.T) {
	for _, test := range []struct {
		desc   string
		ii, jj Interval
		wantAfter,
		wantBefore bool
	}{
		{
			desc: "empty",
		},
		{
			desc:      "after",
			ii:        Interval{9, 12},
			jj:        Interval{5, 8},
			wantAfter: true,
		},
		{
			desc:       "before",
			ii:         Interval{1, 4},
			jj:         Interval{5, 8},
			wantBefore: true,
		},
		{
			desc: "touching after",
			ii:   Interval{8, 9},
			jj:   Interval{5, 8},
		},
		{
			desc: "touching before",
			ii:   Interval{5, 5},
			jj:   Interval{5, 8},
		},
		{
			desc: "including",
			ii:   Interval{1, 9},
			jj:   Interval{5, 8},
		},
		{
			desc: "overlapping",
			ii:   Interval{1, 6},
			jj:   Interval{5, 8},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := test.ii.After(test.jj); got != test.wantAfter {
				t.Errorf("After() = %t, want %t", got, test.wantAfter)
			}
			if got := test.ii.Before(test.jj); got != test.wantBefore {
				t.Errorf("Before() = %t, want %t", got, test.wantBefore)
			}
		})
	}
}
