package network

import (
	"testing"
)

func TestLS(t *testing.T) {
	_, err := LS()
	if err != nil {
		t.Errorf("Error querying for proxy %s", err.Error())
	}
}

func TestInterface(t *testing.T) {
	tests := []struct {
		name     string
		expError error
	}{
		{"lo", nil},
		{"sdadadda", errInterfaceNotFound},
	}

	for _, test := range tests {
		_, err := Interface(test.name)
		if err != test.expError {
			t.Errorf("Expected error to be %v instead got %v", test.expError, err)
		}
	}

}
