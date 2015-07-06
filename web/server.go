package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josselinauguste/sheepit/project"
)

const port = 8080

func main() {
	bus := sheepit.NewSynchronousBus()
	buildResource := newBuildResource(bus)

	r := mux.NewRouter().StrictSlash(false)
	builds := r.Path("/builds").Subrouter()
	builds.Methods("POST").HandlerFunc(buildResource.createBuildHandler)

	fmt.Printf("Sheepit is listening on port %v...", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
