package main

import "fmt"

func main() {

	aa := map[string]*string{}
	aa["example"] = pstring("example")
	fmt.Println(aa)
}

func pstring(v string) *string {
	return &v
}
