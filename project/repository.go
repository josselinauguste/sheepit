package sheepit

import (
	"gopkg.in/libgit2/git2go.v22"
	"io/ioutil"
)

type Repository struct {
	Url  string
	path string
}

func NewRepository(url string) *Repository {
	repository := new(Repository)
	repository.Url = url
	return repository
}

func (r *Repository) Retrieve() error {
	path, err := ioutil.TempDir("", "sheepit")
	if err != nil {
		return err
	}
	_, err = git.Clone(r.Url, path, new(git.CloneOptions))
	if err != nil {
		return err
	}
	r.path = path
	return nil
}

func (r Repository) Path() string {
	return r.path
}
