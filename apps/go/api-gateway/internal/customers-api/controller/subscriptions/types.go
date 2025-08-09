package customerSubscriptionsController

import (
	"net/http"

	"project1.v0/api_gateway/internal/customers/di"
	"project1.v0/pkg/decoder"
	validatorХ "project1.v0/pkg/validator_x"
)

type CustomerSubsciptionsApiControllerDeps struct {
	Router                      *http.ServeMux
	CustomerSubscriptionService di.ICustomerSubscriptionService
	*validatorХ.ValidatorX
	*decoder.DecoderService
}

type CustomerSubsciptionsApiController struct {
	service   di.ICustomerSubscriptionService
	validator *validatorХ.ValidatorX
	decoder   *decoder.DecoderService
}
