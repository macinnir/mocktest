package mocktest

import "fmt"

type SomeDep struct {
}

func (d *SomeDep) SomeMethod() {
	fmt.Println("SomeMethod called")
}

func NewSomeDep() *SomeDep {
	return &SomeDep{}
}
