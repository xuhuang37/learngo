package container

import "fmt"

func printSlice(s []int) {
	fmt.Printf("slice = %v ,len(s)=%d, cap(s)=%d \n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	var s []int //zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println("s = ", s)
	s1 := []int {2,3,4,5}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting element from slice")

	s2 = append(s2[:3] ,s2[4:]... )
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
