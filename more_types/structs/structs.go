package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func MyStruct() {
	
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // same as (*p).X
	fmt.Println(v)
}