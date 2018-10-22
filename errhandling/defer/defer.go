package _defer

import (
	"fmt"
	"learngo/functional/fib"
	"os"
	"bufio"
)

func tryDefer()  {
	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
}
func write(filename string)  {
	file,err:=os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)//create a file
	if err !=nil{
		if pathError,ok := err.(*os.PathError);!ok{
			panic(err)
		}else {
			fmt.Printf("%s ,%s, %s",pathError.Op,pathError.Path,pathError.Err)
		}
	}
	defer file.Close()// close the file

	writer:= bufio.NewWriter(file) //write into the file
	defer writer.Flush()

	f:=fib.Fibonacci()
	for i:=1;i<20;i++{
		fmt.Fprintln(writer,f())
	}
}

func main() {
	//tryDefer()
	write("/fibonacci.txt")
}

