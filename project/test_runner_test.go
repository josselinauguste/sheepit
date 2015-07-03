package sheepit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunProjectTests(t *testing.T) {
	repository := NewRepository(lookupGitRepository(t, "basic_repository"))
	repository.Retrieve()

	log, ok, err := RunTests(*repository)

	assert.True(t, ok)
	assert.NotEmpty(t, log)
	assert.Nil(t, err)
}

func TestGetErrorIfNoCommandFound(t *testing.T) {
	repository := NewRepository("/")

	_, ok, err := RunTests(*repository)

	assert.False(t, ok)
	assert.NotNil(t, err)
}
