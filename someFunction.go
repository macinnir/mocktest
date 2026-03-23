package mocktest

import "github.com/macinnir/mocktest/subpackage"

// Inject some dependencies and call their methods
func SomeFunction(
	dep SomeDepInterface,
	dep2 AnotherDepInterface,
	subDep subpackage.SomeSubDepInterface,
) string {
	return dep.SomeMethod() + dep2.AnotherMethod() + subDep.SomeSubMethod()
}
