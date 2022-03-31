package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = "100:200"
	var s2 = ":200"
	var s3 = "100:"
	s2 = strings.TrimSuffix(s2, ":")
	s3 = strings.TrimSuffix(s3, ":")
	fmt.Println(strings.Split(s1, ":"))
	fmt.Println(strings.Split(s2, ":"))
	fmt.Println(strings.Split(s3, ":"))
}