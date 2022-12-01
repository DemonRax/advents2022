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
			in: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			want: 24000,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := maxCalories(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
