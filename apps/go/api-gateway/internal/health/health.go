package health

import (
	"fmt"
	"net/http"
	"time"

	mdw "project1.v0/pkg/http/middleware"
	"project1.v0/pkg/http/resp"
)

type HealthApiModuleDeps struct {
	Router *http.ServeMux
}

type HealthApiModule struct {
}

type HealthApiController struct {
}

func NewHealthApiModule(deps HealthApiModuleDeps) *HealthApiModule {

	///////////////
	// Controllers
	NewHealthApiController(HealthApiControllerDeps{Router: deps.Router})

	return &HealthApiModule{}
}

type HealthApiModuleConstructor func(deps HealthApiModuleDeps) *HealthApiModule

type HealthApiControllerDeps struct {
	Router *http.ServeMux
}

func NewHealthApiController(deps HealthApiControllerDeps) {

	handler := HealthApiController{}

	///////////
	//WEB
	deps.Router.Handle("GET /health", mdw.Chain()(http.HandlerFunc(handler.Health())))

}

func (h *HealthApiController) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp.Json(w, fmt.Sprintf("OK, now is %v", time.Now()), 200)

	}
}
