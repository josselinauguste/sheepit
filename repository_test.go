package sheepit

import (
	"io/ioutil"
	"testing"
)

func TestNewRepository(t *testing.T) {
	url := "/tmp/fake"

	repository := NewRepository(url)

	if repository == nil {
		t.Error("Expected not nil repository")
	}
	if repository.Url != url {
		t.Errorf("Expected %v, got %v", url, repository.Url)
	}
}

func TestGetRepository(t *testing.T) {
	url := createTestGitRepository(t)
	repository := NewRepository(url)

	repository.Retrieve()

	if repository.Path == "" {
		t.Error("Expected not nil path")
	}
	files, _ := ioutil.ReadDir(repository.Path)
	if len(files) != 2 {
		t.Error("Expected only one file")
	}
	if files[1].Name() != "README" {
		t.Errorf("Expected to be README, got %v", files[0].Name())
	}
}
