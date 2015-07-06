package sheepit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FakeCommand struct {
	mock.Mock
}

func (m *FakeCommand) execute() error {
	m.Called()
	return nil
}

func TestSendToSynchronousBus(t *testing.T) {
	bus := NewSynchronousBus()
	command := new(FakeCommand)
	command.On("execute").Return()

	err := bus.Send(command)

	command.AssertExpectations(t)
	assert.Nil(t, err)
}
