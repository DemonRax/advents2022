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
			in:   `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
			want: 19,
		},
		{
			in:   `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			want: 23,
		},
		{
			in:   `nppdvjthqldpwncqszvftbrmjlhg`,
			want: 23,
		},
		{
			in:   `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			want: 29,
		},
		{
			in:   `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
			want: 26,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := firstUnique(util.ReadString(test.in)[0], 14); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
