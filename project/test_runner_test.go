package sheepit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunProjectTests(t *testing.T) {
	repository := NewRepository(lookupGitRepository("basic_repository"))
	repository.Retrieve()

	log, ok := RunTests(*repository)

	assert.True(t, ok)
	assert.NotEmpty(t, log)
}
