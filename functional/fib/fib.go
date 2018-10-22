package fib

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)


type getInt func() int

//给函数实现接口

func (g getInt) Read(p []byte) (n int, err error) {
	next:=g()
	if next>=10000{
		return 0,io.EOF
	}
	s:=fmt.Sprintf("%d\n",next)
	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



func Fibonacci() getInt  {
	a,b:=0,1
	fmt.Println(b)
	return func() int {
		a,b = b, a+b
		return b
	}

}
//func main() {
//	a:=fibonacci()
//	printFileContents(a)
//}
