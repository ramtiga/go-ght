package main

import (
	"fmt"
	"github.com/ramtiga/go-ght"
)

func main() {
	repo := ght.GetRepoInf("go")

	for i, r := range repo {
		fmt.Printf("%d : %s\n", i+1, r.Name)
	}
}
