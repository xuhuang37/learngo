package container

import "fmt"

//import "fmt"

func lengthOfNoneRepeatingSubStr(s string) int {
	m := make(map[rune]int)
	start := 0
	length:=0
	for i, ch := range []rune(s) {
		if lastI,ok:=m[ch]; ok&&lastI >= start{
			start = m[ch]+ 1
		}
		if i-start+1>length{
			length = i-start+1
		}
		m[ch] = i
	}
	fmt.Println(length)
	return length
}

func main() {
	lengthOfNoneRepeatingSubStr("我是是个好人")
}
