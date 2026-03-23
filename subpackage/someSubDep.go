package subpackage

// Implements SomeSubDepInterface
type SomeSubDep struct {
}

func (s *SomeSubDep) SomeSubMethod() string {
	return "subpackage.SomeSubMethod() called"
}

func NewSomeSubDep() *SomeSubDep {
	return &SomeSubDep{}
}
