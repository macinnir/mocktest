package mocktest

type SomeDep struct {
}

func (d *SomeDep) SomeMethod() string {
	return "SomeDep.SomeMethod() called"
}

func NewSomeDep() *SomeDep {
	return &SomeDep{}
}
