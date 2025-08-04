package customerSubscriptionsRepository

import (
	customerSubscriptionModel "project1.v0/api_gateway/internal/customers/models/subscription"
	dbMapper "project1.v0/mappers/dto_db"
	db "project1.v0/pkg/db/postgresql"
)

type CustomerSubscriptionRepositoryDeps struct {
	Db                 *db.Db
	DtosToDbMapService *dbMapper.DtosToDbMapService
}

type CustomerSubscriptionRepository struct {
	db                 *db.Db
	DtosToDbMapService *dbMapper.DtosToDbMapService
}

func NewCustomerSubscriptionRepository(deps CustomerSubscriptionRepositoryDeps) *CustomerSubscriptionRepository {
	deps.DtosToDbMapService.AddMapping(&customerSubscriptionModel.CustomerSubscription{})

	return &CustomerSubscriptionRepository{
		db:                 deps.Db,
		DtosToDbMapService: deps.DtosToDbMapService,
	}
}
