package customerSubscriptionsRepository

import (
	//customerSubscriptionModel "project1.v0/api_gateway/internal/customers/models/subscription"
	//dbMapper "project1.v0/mappers/dto_db"
	db "project1.v0/pkg/db/postgresql"
	"project1.v0/pkg/decoder"
)

type CustomerSubscriptionRepositoryDeps struct {
	Db *db.Db
	//DtosToDbMapService *dbMapper.DtosToDbMapService
	*decoder.DecoderService
}

type CustomerSubscriptionRepository struct {
	db *db.Db
	//DtosToDbMapService *dbMapper.DtosToDbMapService
	decoder *decoder.DecoderService
}

func NewCustomerSubscriptionRepository(deps CustomerSubscriptionRepositoryDeps) *CustomerSubscriptionRepository {
	//deps.DtosToDbMapService.AddMapping(&customerSubscriptionModel.CustomerSubscription{})

	return &CustomerSubscriptionRepository{
		db: deps.Db,
		//DtosToDbMapService: deps.DtosToDbMapService,
		decoder: deps.DecoderService,
	}
}
