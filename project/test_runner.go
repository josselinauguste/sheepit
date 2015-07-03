package sheepit

import (
	"errors"
	"os"
	"os/exec"
	"path"
)

func RunTests(project Project) ([]byte, bool, error) {
	command := getRunnerCommand(project.Path())
	if command == nil {
		return nil, false, errors.New("Can't find a valid command to run tests")
	}
	output, err := command.Output()
	return output, err == nil, nil
}

func getRunnerCommand(projectPath string) *exec.Cmd {
	var command *exec.Cmd
	if _, err := os.Stat(path.Join(projectPath, "run_tests.sh")); err == nil {
		command = exec.Command("./run_tests.sh")
	} else if _, err := os.Stat(path.Join(projectPath, "Rakefile")); err == nil {
		command = exec.Command("rake", "test")
	}
	if command != nil {
		command.Dir = projectPath
	}
	return command
}
