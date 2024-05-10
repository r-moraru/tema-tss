package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerIsInitializedCorrectly(t *testing.T) {
	_, stopServer, err := RunRpcServer(0)
	defer stopServer()

	assert.NoError(t, err, "Server initialization should not return an error")
}
