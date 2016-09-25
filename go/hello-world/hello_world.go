package hello

import "fmt"

const testVersion = 2

func HelloWorld(name string) string {
	fmt.Println(name)
	if len(name) == 0 {
		name = "World"
	}
	return "Hello, " + name + "!"
}
