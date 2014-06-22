package main

import (
	"fmt"
	"github.com/ramtiga/go-ght"
)

func main() {
	repo := ght.GetRepoInf("go")

	for i, r := range repo {
		fmt.Printf("%3d  %-40s  %4s %4s\n", i+1, r.Name, r.Star, r.Fork)
	}
}
