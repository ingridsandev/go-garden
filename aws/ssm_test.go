package aws

import (
	"testing"
)

const ParameterName string = "/go/garden/chrysanthemum/test"
const ParameterValue string = "Go and AWS are amazing"

func TestPutParameter(t *testing.T) {
	want := "OK"
	if got := PutParameter("eu-west-1", ParameterName, ParameterValue); got != want {
		t.Errorf("SSM.PutParameter() = %q, want %q", got, want)
	}
}

func TestGetParameter(t *testing.T) {
	if got := GetParameter("eu-west-1", ParameterName); got != ParameterValue {
		t.Errorf("SSM.PutParameter() = %q, want %q", got, ParameterValue)
	}
}

func TestDeleteParameter(t *testing.T) {
	want := "OK"
	if got := DeleteParameter("eu-west-1", ParameterName); got != want {
		t.Errorf("SSM.PutParameter() = %q, want %q", got, want)
	}
}