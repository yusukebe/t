package t

import (
	"testing"
)

func TestTest(testing *testing.T) {
	inputs := []struct {
		expected string
		param    param
		received string
	}{
		{expected: "", param: param{}, received: ""},
		{expected: "hello", param: param{}, received: "hello"},
		{expected: "2", param: param{operator: "==", evalFlag: true}, received: "1+1"},
		{expected: "10", param: param{operator: "<", evalFlag: true}, received: "2*10"},
		{expected: "a.*", param: param{operator: "=~", matchFlag: true}, received: "abc"},
	}

	for _, input := range inputs {
		passed := test(input.expected, input.received, input.param)
		if passed == false {
			testing.Errorf("%v %s %v: is not true", input.expected, input.param.operator, input.received)
		}
	}

	inputs = []struct {
		expected string
		param    param
		received string
	}{
		{expected: "", param: param{operator: "=="}, received: "h"},
		{expected: "hello", param: param{operator: "=="}, received: "hell"},
	}

	for _, input := range inputs {
		passed := test(input.expected, input.received, input.param)
		if passed == true {
			testing.Errorf("%v is %v", input.expected, input.received)
		}
	}
}
