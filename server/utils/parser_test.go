package utils

import (
	"testing"
)
// TODO add more tests
func TestGetOrderId(t *testing.T) {
	expectedResult := `123abc`
	actualResult, err := GetOrderId(`/prefix/123abc`, `prefix`)
	if expectedResult != actualResult || err != nil {
		t.Fatalf(`GetOrderId("/prefix/123abc"", "prefix"") = %q, %v, want expected %#q, nil`, actualResult, err, expectedResult)
	}
}
