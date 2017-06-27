package main

import "sort"
import "fmt"

type people []string

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
}
