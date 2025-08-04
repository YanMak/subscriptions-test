package customerApiModule

import (
	"net/http"

	customerSubscriptionsController "project1.v0/api_gateway/internal/customers-api/controller/subscriptions"
	"project1.v0/api_gateway/internal/customers/di"

	//"project1.v0/configs"
	validatorХ "project1.v0/pkg/validator_x"
)

type CustomerApiModuleDeps struct {
	Router *http.ServeMux
	//*customerSubscriptionService.CustomerSubscriptionService
	CustomerSubscriptionService di.ICustomerSubscriptionService
	*validatorХ.ValidatorX
}

type CustomerApiModule struct {
}

func NewCustomerApiModule(deps CustomerApiModuleDeps) *CustomerApiModule {

	////////////
	// Services

	///////////////
	// Controllers
	customerSubscriptionsController.NewCustomerSubsciptionsApiController(
		customerSubscriptionsController.CustomerSubsciptionsApiControllerDeps{
			Router:                      deps.Router,
			CustomerSubscriptionService: deps.CustomerSubscriptionService,
			ValidatorX:                  deps.ValidatorX,
		},
	)

	return &CustomerApiModule{}
}

type CustomerApiModuleConstructor func(deps CustomerApiModuleDeps) *CustomerApiModule
