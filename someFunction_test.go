package mocktest_test

import (
	"testing"

	"github.com/macinnir/mocktest"
	"github.com/macinnir/mocktest/subpackage"
	"github.com/stretchr/testify/assert"
)

func TestSomeFunction(t *testing.T) {
	// Create mocks for the dependencies
	mockDep := mocktest.NewMockSomeDepInterface(t)
	mockDep2 := mocktest.NewMockAnotherDepInterface(t)
	mockSubDep := subpackage.NewMockSomeSubDepInterface(t)

	// Set up expectations for the mocks
	mockDep.EXPECT().SomeMethod().Return("mocked SomeMethod() called")
	mockDep2.EXPECT().AnotherMethod().Return("mocked AnotherMethod() called")
	mockSubDep.EXPECT().SomeSubMethod().Return("mocked SomeSubMethod() called")

	// Call the function under test with the mocked dependencies
	result := mocktest.SomeFunction(mockDep, mockDep2, mockSubDep)

	// Assert the expected result
	expected := "mocked SomeMethod() calledmocked AnotherMethod() calledmocked SomeSubMethod() called"
	assert.Equal(t, expected, result)
}
