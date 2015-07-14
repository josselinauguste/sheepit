package web

import (
	"net/http"
	"strings"
	"testing"

	"github.com/josselinauguste/magicbus"
	"github.com/josselinauguste/sheepit/project"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FakeBus struct {
	mock.Mock
}

func (m *FakeBus) Send(command magicbus.Command) error {
	m.Called(command)
	return nil
}

func TestCreateBuild(t *testing.T) {
	fakeBus := new(FakeBus)
	fakeBus.On("Send", sheepit.NewCreateBuildCommand("http://git.com")).Return(nil)
	resource := newBuildResource(fakeBus)
	request, _ := http.NewRequest("POST", "/builds", strings.NewReader(`{"url": "http://git.com"}`))
	request.Header.Set("Content-Type", "application/json")
	response := NewFakeResponse(t)

	resource.createBuildHandler(response, request)

	response.AssertStatus(http.StatusOK)
	fakeBus.AssertExpectations(t)
	buildCreated := &buildCreated{}
	err := response.GetJsonBody(buildCreated)
	assert.Nil(t, err)
	assert.True(t, buildCreated.Success)
	assert.Empty(t, buildCreated.Output)
}
