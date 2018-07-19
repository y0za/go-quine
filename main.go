package main

import "fmt"

func main() {
	s := `package main

import "fmt"

func main() {
	s := %[2]s%[1]s%[2]s
	fmt.Printf(s, s, "%[2]s")
}
`
	fmt.Printf(s, s, "`")
}
