package t

import (
	"testing"
)

func TestTest(testing *testing.T) {
	inputs := []struct {
		expected string
		param    param
		actual   string
	}{
		{expected: "", param: param{}, actual: ""},
		{expected: "hello", param: param{}, actual: "hello"},
		{expected: "2", param: param{operator: "==", evalFlag: true}, actual: "1+1"},
		{expected: "10", param: param{operator: "<", evalFlag: true}, actual: "2*10"},
	}

	for _, input := range inputs {
		passed := test(input.expected, input.actual, input.param)
		if passed == false {
			testing.Errorf("%v %s %v: is not true", input.expected, input.param.operator, input.actual)
		}
	}

	inputs = []struct {
		expected string
		param    param
		actual   string
	}{
		{expected: "", param: param{operator: "=="}, actual: "h"},
		{expected: "hello", param: param{operator: "=="}, actual: "hell"},
	}

	for _, input := range inputs {
		passed := test(input.expected, input.actual, input.param)
		if passed == true {
			testing.Errorf("%v is %v", input.expected, input.actual)
		}
	}
}
