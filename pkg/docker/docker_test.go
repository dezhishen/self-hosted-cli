package docker

import "testing"

func TestStartContainer(t *testing.T) {
	StartContainer("library/alpine", []string{"sleep", "30"}, nil, nil, nil)
}
