package container

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网！"
	for _, v := range s {
		fmt.Printf("%X ",v)
	}
	fmt.Println()
	for i, v := range s {
		fmt.Printf("(%d %X)", i, v)
	}
	fmt.Println()
	fmt.Println("rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes)>0{
	ch,size := utf8.DecodeRune(bytes)
	bytes = bytes[size:]
	fmt.Printf("%c  ",ch)
	}
	fmt.Println()
	for i,v := range []rune(s){
		fmt.Printf("(%d,%c)",i,v)
	}

}
