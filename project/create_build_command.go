package sheepit

import "fmt"

type FailingTestError struct {
	output string
}

func (e FailingTestError) Error() string {
	return fmt.Sprintf("Failing test:\n%v", e.output)
}

type CreateBuildCommand struct {
	url string
}

func NewCreateBuildCommand(url string) *CreateBuildCommand {
	return &CreateBuildCommand{url}
}

func (command CreateBuildCommand) Execute() error {
	repository := NewRepository(command.url)
	err := repository.Retrieve()
	if err != nil {
		return err
	}
	output, success, err := RunTests(repository)
	if err != nil {
		return err
	}
	if !success {
		return FailingTestError{string(output)}
	}
	return nil
}
