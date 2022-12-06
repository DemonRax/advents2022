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
			want: 7,
		},
		{
			in:   `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			want: 5,
		},
		{
			in:   `nppdvjthqldpwncqszvftbrmjlhg`,
			want: 6,
		},
		{
			in:   `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			want: 10,
		},
		{
			in:   `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
			want: 11,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := firstUnique(util.ReadString(test.in)[0], 4); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
