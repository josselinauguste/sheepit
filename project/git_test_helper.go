package sheepit

import (
	"github.com/termie/go-shutil"
	"gopkg.in/libgit2/git2go.v22"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func lookupGitRepository(name string) string {
	sourcePath := path.Join(lookupRoot(), "fixtures", name)
	return instantiateGitRepository(sourcePath)
}

func instantiateGitRepository(sourcePath string) string {
	destPath, _ := ioutil.TempDir("", "sheepit_test")
	destPath = path.Join(destPath, "repository")
	shutil.CopyTree(sourcePath, destPath, nil)
	initGitRepository(destPath)
	return destPath
}

func initGitRepository(path string) {
	repository, _ := git.InitRepository(path, false)
	loc, _ := time.LoadLocation("Europe/Berlin")
	sig := &git.Signature{
		Name:  "Rand Om Hacker",
		Email: "random@hacker.com",
		When:  time.Date(2013, 03, 06, 14, 30, 0, 0, loc),
	}

	idx, _ := repository.Index()
	_ =
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			idx.AddByPath(info.Name())
			return nil
		})
	treeId, _ := idx.WriteTree()

	message := "Commit\n"
	tree, _ := repository.LookupTree(treeId)
	repository.CreateCommit("HEAD", sig, sig, message, tree)
}

func lookupRoot() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
