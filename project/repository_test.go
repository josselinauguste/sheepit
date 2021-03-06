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
	url := lookupGitRepository(t, "basic_repository")
	repository := NewRepository(url)

	err := repository.Retrieve()

	assert.Nil(t, err)
	assert.NotEmpty(t, repository.Path())
	files, _ := ioutil.ReadDir(repository.Path())
	assert.Equal(t, 3, len(files))
	assert.Equal(t, "README", files[1].Name())
}
