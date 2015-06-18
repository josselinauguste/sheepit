package sheepit

import (
	"os"
	"os/exec"
	"path"
)

func RunTests(project Project) ([]byte, bool) {
	command := getRunnerCommand(project.Path())
	output, err := command.Output()
	if err != nil {
		return output, false
	}
	return output, true
}

func getRunnerCommand(projectPath string) *exec.Cmd {
	var command *exec.Cmd
	if _, err := os.Stat(path.Join(projectPath, "run_tests.sh")); err == nil {
		command = exec.Command("./run_tests.sh")
	} else if _, err := os.Stat(path.Join(projectPath, "Rakefile")); err == nil {
		command = exec.Command("rake", "test")
	}
	command.Dir = projectPath
	return command
}
