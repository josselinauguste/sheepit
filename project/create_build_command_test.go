package sheepit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCreateBuildCommand(t *testing.T) {
	url := lookupGitRepository(t, "basic_repository")
	command := NewCreateBuildCommand(url)

	err := command.Execute()

	assert.Nil(t, err)
}

func TestExecuteCreateFailingBuildCommand(t *testing.T) {
	url := lookupGitRepository(t, "failing_repository")
	command := NewCreateBuildCommand(url)

	err := command.Execute()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Fake output")
}
