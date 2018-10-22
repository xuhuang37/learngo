package container

import "testing"

func TestNoneRepeating(t *testing.T)  {
	cases := []struct{
		s string
		ans int
	}{

		{"abcabcbb", 3},
		{"pwwkew", 3},
		// edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},
		//chinese support
		{"这里是慕课网", 6},
		{"一二三三二一", 3},
	}
	for _,ts := range cases{
		actual := lengthOfNoneRepeatingSubStr(ts.s)
		if ts.ans!= actual{
			t.Errorf("got %d for int %s;expected:%d",ts.ans,ts.s, actual )
		}
	}

}
