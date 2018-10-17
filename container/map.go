package container

import "fmt"

func main() {
	m:=map[string]string{
		"name":"hx",
		"course":"golang",
		"site":"home",
		"quality":"notbad",
	}

	m2 := make(map[string]int)// empty map

	var m3 map[string]int // nil

	fmt.Println(m,m2,m3)

	fmt.Println("Traversing map")

	for k,v:= range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	fmt.Println("spelling error")
	causeName,ok := m["cause"]
	fmt.Println(causeName,ok) //spelling error make a zero element

	if causeName,ok :=m["cause"];ok{
		fmt.Println(causeName)
	}else{
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name,ok:=m["name"]
	fmt.Println(name,ok)
	delete(m, "name")
	name,ok=m["name"]
	fmt.Println(name,ok)
}
