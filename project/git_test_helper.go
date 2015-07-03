package sheepit

import (
	"github.com/termie/go-shutil"
	"gopkg.in/libgit2/git2go.v22"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

//TODO instantiate only once always in the same directory

func lookupGitRepository(t *testing.T, name string) string {
	sourcePath := path.Join(lookupRoot(), "fixtures", name)
	return instantiateGitRepository(t, sourcePath)
}

func instantiateGitRepository(t *testing.T, sourcePath string) string {
	destPath, err := ioutil.TempDir("", "sheepit_test")
	if err != nil {
		t.Error("Can't create git repository instance")
	}
	destPath = path.Join(destPath, "repository")
	shutil.CopyTree(sourcePath, destPath, nil)
	initGitRepository(t, destPath)
	return destPath
}

func initGitRepository(t *testing.T, path string) {
	repository, err := git.InitRepository(path, false)
	if err != nil {
		t.Error("Can't init git repository")
	}
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Error("Can't load location")
	}
	sig := &git.Signature{
		Name:  "Rand Om Hacker",
		Email: "random@hacker.com",
		When:  time.Date(2013, 03, 06, 14, 30, 0, 0, loc),
	}

	idx, err := repository.Index()
	if err != nil {
		t.Error("Can't index git repository")
	}
	_ =
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			idx.AddByPath(info.Name())
			return nil
		})
	treeId, err := idx.WriteTree()
	if err != nil {
		t.Error("Can't write tree")
	}

	message := "Commit\n"
	tree, err := repository.LookupTree(treeId)
	if err != nil {
		t.Error("Can't lookup tree")
	}
	repository.CreateCommit("HEAD", sig, sig, message, tree)
}

func lookupRoot() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
