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
			want: 8,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := highestScore(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_score(t *testing.T) {
	grid := generateGrid(util.ReadString(`30373
25512
65332
33549
35390`))

	for _, test := range []struct {
		desc string
		tree,
		i0, iN, j0, jN int
		want int
	}{
		{
			desc: "up 1 2",
			tree: 5,
			i0:   0,
			iN:   0,
			j0:   2,
			jN:   2,
			want: 1,
		},
		{
			desc: "down 1 2",
			tree: 5,
			i0:   2,
			iN:   4,
			j0:   2,
			jN:   2,
			want: 2,
		},
		{
			desc: "left 1 2",
			tree: 5,
			i0:   1,
			iN:   1,
			j0:   1,
			jN:   0,
			want: 1,
		},
		{
			desc: "right 1 2",
			tree: 5,
			i0:   1,
			iN:   1,
			j0:   3,
			jN:   4,
			want: 2,
		},
		{
			desc: "up 3 2",
			tree: 5,
			i0:   0,
			iN:   2,
			j0:   2,
			jN:   2,
			want: 2,
		},
		{
			desc: "down 3 2",
			tree: 5,
			i0:   4,
			iN:   4,
			j0:   2,
			jN:   2,
			want: 1,
		},
		{
			desc: "left 3 2",
			tree: 5,
			i0:   3,
			iN:   3,
			j0:   1,
			jN:   0,
			want: 2,
		},
		{
			desc: "right 3 2",
			tree: 5,
			i0:   3,
			iN:   3,
			j0:   3,
			jN:   4,
			want: 2,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := score(test.tree, test.i0, test.iN, test.j0, test.jN, grid); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
