package main

import (
	"testing"

	"advents2022/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `30373
25512
65332
33549
35390`,
			want: 21,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := visibleTrees(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_visibleRow(t *testing.T) {
	for _, test := range []struct {
		desc string
		j    int
		want bool
	}{
		{
			desc: "visible edge left",
			j:    0,
			want: true,
		},
		{
			desc: "visible left through right",
			j:    1,
			want: true,
		},
		{
			desc: "not visible center",
			j:    2,
			want: false,
		},
		{
			desc: "visible right through right",
			j:    3,
			want: true,
		},
		{
			desc: "visible edge right",
			j:    4,
			want: true,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := visibleRow(0, test.j, [][]int{{6, 5, 3, 3, 2}}); got != test.want {
				t.Errorf("want %t got %t", test.want, got)
			}
		})
	}
}

func Test_visibleCol(t *testing.T) {
	for _, test := range []struct {
		desc string
		i    int
		want bool
	}{
		{
			desc: "visible top",
			i:    0,
			want: true,
		},
		{
			desc: "visible from top",
			i:    1,
			want: true,
		},
		{
			desc: "not visible",
			i:    2,
			want: false,
		},
		{
			desc: "not visible",
			i:    3,
			want: false,
		},
		{
			desc: "visible bottom",
			i:    4,
			want: true,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := visibleCol(test.i, 0, [][]int{{0}, {5}, {5}, {3}, {5}}); got != test.want {
				t.Errorf("want %t got %t", test.want, got)
			}
		})
	}
}
