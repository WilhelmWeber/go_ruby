package test

import "fmt"

func main() {
	var env map[string]int
	list := []string{"aiu", "eio", "kaki", "kuke", "koki"}

	for _, e := range list {
		test(e, env)
	}
}

func test(s string, env map[string]int) {
	fmt.Println(env)
	env[s] = 1
}
