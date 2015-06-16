package sheepit

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewRepository(t *testing.T) {
	url := "/tmp/fake"

	repository := NewRepository(url)

	assert.NotNil(t, repository)
	assert.Equal(t, url, repository.Url)
}

func TestGetRepository(t *testing.T) {
	url := createTestGitRepository(t)
	repository := NewRepository(url)

	repository.Retrieve()

	assert.NotEmpty(t, repository.Path)
	files, _ := ioutil.ReadDir(repository.Path)
	assert.Equal(t, 2, len(files))
	assert.Equal(t, "README", files[1].Name())
}
