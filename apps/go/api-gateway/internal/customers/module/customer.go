package customerModule

import (
	"project1.v0/api_gateway/internal/customers/di"
	customerSubscriptionsRepository "project1.v0/api_gateway/internal/customers/repository/subscription"
	customerSubscriptionService "project1.v0/api_gateway/internal/customers/service/subscription"
	db "project1.v0/pkg/db/postgresql"
	"project1.v0/pkg/decoder"
	validatorХ "project1.v0/pkg/validator_x"
)

type CustomerModuleDeps struct {
	*db.Db
	*validatorХ.ValidatorX
	*decoder.DecoderService
}

type CustomerModule struct {
	//*customerSubscriptionService.CustomerSubscriptionService
	CustomerSubscriptionService di.ICustomerSubscriptionService
}

func NewCustomerModule(deps CustomerModuleDeps) *CustomerModule {

	repo := customerSubscriptionsRepository.NewCustomerSubscriptionRepository(customerSubscriptionsRepository.CustomerSubscriptionRepositoryDeps{
		Db: deps.Db,
		//DtosToDbMapService: deps.DtosToDbMapService,
	})

	service := customerSubscriptionService.NewCustomerSubscriptionService(customerSubscriptionService.CustomerSubscriptionServiceDeps{
		Repo: repo,
	})
	return &CustomerModule{CustomerSubscriptionService: service}
}

type CustomerModuleConstructor func(deps CustomerModuleDeps) *CustomerModule
