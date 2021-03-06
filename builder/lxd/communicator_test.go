package lxd

import (
	"github.com/hashicorp/packer/packer"
	"testing"
)

func TestCommunicator_ImplementsCommunicator(t *testing.T) {
	var raw interface{}
	raw = &Communicator{}
	if _, ok := raw.(packer.Communicator); !ok {
		t.Fatalf("Communicator should be a communicator")
	}
}

// Acceptance tests
// TODO Execute a command
// TODO Upload a file
// TODO Download a file
// TODO Upload a Directory
