package web

import (
	"encoding/json"
	"net/http"

	"github.com/josselinauguste/magicbus"
	"github.com/josselinauguste/sheepit/project"
)

type buildResource struct {
	bus magicbus.Bus
}

type createBuild struct {
	Url string
}

type buildCreated struct {
	Success bool
	Output  string
}

func newBuildResource(bus magicbus.Bus) *buildResource {
	return &buildResource{bus}
}

func (resource buildResource) createBuildHandler(rw http.ResponseWriter, r *http.Request) {
	createBuild := new(createBuild)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(createBuild)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	command := sheepit.NewCreateBuildCommand(createBuild.Url)
	err = resource.bus.Send(command)

	buildCreated := new(buildCreated)
	buildCreated.Success = err == nil
	if err != nil {
		buildCreated.Output = err.Error()
	}
	response, err := json.Marshal(buildCreated)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(response)
}
