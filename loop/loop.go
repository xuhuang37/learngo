package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
	s := `	aaaaa
			bbbbb
			*******
			77777
	`
	printFile("abc.txt")
	printFileContents(strings.NewReader(s))
}
