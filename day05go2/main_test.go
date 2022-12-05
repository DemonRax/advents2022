package main

import (
	"testing"

	"advents2022/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in,
		want string
	}{
		{
			in: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`,
			want: "MCD",
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := top(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %s, want %s", got, test.want)
			}
		})
	}
}
