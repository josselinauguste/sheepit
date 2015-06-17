package sheepit

import (
	"os/exec"
)

func RunTests(project Project) ([]byte, bool) {
	command := exec.Command("./run_tests.sh")
	command.Dir = project.Path()
	output, err := command.Output()
	if err != nil {
		return output, false
	}
	return output, true
}
