package sheepit

import (
	"gopkg.in/libgit2/git2go.v22"
	"io/ioutil"
	"testing"
	"time"
)

func createTestGitRepository(t *testing.T) string {
	path, err := ioutil.TempDir("", "git2go")
	if err != nil {
		t.Error("Can't create test git repository folder")
	}
	repository, err := git.InitRepository(path, false)
	if err != nil {
		t.Error("Can't init test git repository")
	}

	seedTestGitRepository(t, repository, path)

	return path
}

func seedTestGitRepository(t *testing.T, repo *git.Repository, path string) {
	ioutil.WriteFile(path+"/README", []byte("foo\n"), 0644)

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Error("Expected to not be nil")
	}
	sig := &git.Signature{
		Name:  "Rand Om Hacker",
		Email: "random@hacker.com",
		When:  time.Date(2013, 03, 06, 14, 30, 0, 0, loc),
	}

	idx, err := repo.Index()
	if err != nil {
		t.Error("Expected to not be nil")
	}
	err = idx.AddByPath("README")
	if err != nil {
		t.Error("Expected to not be nil")
	}
	treeId, err := idx.WriteTree()
	if err != nil {
		t.Error("Expected to not be nil")
	}

	message := "This is a commit\n"
	tree, err := repo.LookupTree(treeId)
	if err != nil {
		t.Error("Expected to not be nil")
	}
	_, err = repo.CreateCommit("HEAD", sig, sig, message, tree)
	if err != nil {
		t.Error("Expected to not be nil")
	}
}
