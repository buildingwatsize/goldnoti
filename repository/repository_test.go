package repository_test

import (
	"goldnoti/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToFloat64_Input_Empty_Output_Should_be_Zero(t *testing.T) {
	input := ""
	expected := 0.00

	actual := repository.ToFloat64(input)

	assert.Equal(t, expected, actual, "Expecting result should be zero")
}

func Test_ToFloat64_Input_String_100_Output_Should_be_100(t *testing.T) {
	input := "100"
	expected := 100.00

	actual := repository.ToFloat64(input)

	assert.Equal(t, expected, actual, "Expecting result should be 100.00")
}

func Test_ToFloat64_Input_String_With_Comma_100000_Output_Should_be_100(t *testing.T) {
	input := "100,000"
	expected := 100000.00

	actual := repository.ToFloat64(input)

	assert.Equal(t, expected, actual, "Expecting result should be 100.00")
}

func Test_ToFloat64_Input_String_With_Comma_10000000_Output_Should_be_100(t *testing.T) {
	input := "10,000,000"
	expected := 10000000.00

	actual := repository.ToFloat64(input)

	assert.Equal(t, expected, actual, "Expecting result should be 100.00")
}

func Test_ToFloat64_Input_String_Without_Comma_100000_Output_Should_be_100(t *testing.T) {
	input := "100000"
	expected := 100000.00

	actual := repository.ToFloat64(input)

	assert.Equal(t, expected, actual, "Expecting result should be 100.00")
}
