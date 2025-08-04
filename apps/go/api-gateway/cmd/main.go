// @title           API-Gateway
// @version         1.0
// @description     REST service of online subscriptions at present moment ()
// @contact.name    Yan
// @contact.email   89269167763@yandex.ru.com
package main

import (
	"fmt"
	"net/http"

	"project1.v0/api_gateway/configs"
	db "project1.v0/pkg/db/postgresql"
	httpMdw "project1.v0/pkg/http/middleware"

	customerApiModule "project1.v0/api_gateway/internal/customers-api/module"
	customerModule "project1.v0/api_gateway/internal/customers/module"

	dbMapper "project1.v0/mappers/dto_db"
	validatorХ "project1.v0/pkg/validator_x"
)

type AppDeps struct {
	*configs.Config
	customerModule.CustomerModuleConstructor
	customerApiModule.CustomerApiModuleConstructor
	db.NewDbConstructor
	validatorХ.ValidatorXModuleConstructor
}

type AppServicesType = struct {
	// OtpService            *otp.OtpService
	// KafkaRPCBrokerService *kafka_v3.KafkaRPCBrokerService
	// LoggerService         *loggerLib.LoggerService
	// GrpcService           *grpcLib.GrpcService
}

type AppReturningValue struct {
	HttpHandler http.Handler
	Services    AppServicesType
}

type App func(*AppDeps) *AppReturningValue

func main() {
	conf := configs.LoadConfig()

	app := NewApp(AppDeps{
		Config:                       conf,
		CustomerModuleConstructor:    customerModule.NewCustomerModule,
		CustomerApiModuleConstructor: customerApiModule.NewCustomerApiModule,
		NewDbConstructor:             db.NewDb,
		ValidatorXModuleConstructor:  validatorХ.NewValidatorXModule,
	})

	server := http.Server{
		//Addr: ":8081",
		//Addr: ":3999",
		Addr: ":4011",
		//Handler: middleware.CORS(middleware.Logging(router)),
		Handler: app.HttpHandler,
	}

	fmt.Println("Server is listening on port ", server.Addr)
	server.ListenAndServe()
}

func NewApp(deps AppDeps) *AppReturningValue {

	db := deps.NewDbConstructor(db.DbDeps{
		DbConfig: &deps.Db,
	})
	dtoToDbMapperService := dbMapper.NewDtosToDbMapService()
	validatorX := deps.ValidatorXModuleConstructor(validatorХ.ValidatorXModuleDeps{}).ValidatorX

	router := http.NewServeMux()

	stack := httpMdw.Chain(
	// cors, logging, etc
	)

	customerModule := deps.CustomerModuleConstructor(customerModule.CustomerModuleDeps{
		Db:                 db,
		DtosToDbMapService: dtoToDbMapperService,
		ValidatorX:         validatorX,
	})
	customerSubscriptionService := customerModule.CustomerSubscriptionService

	deps.CustomerApiModuleConstructor(
		customerApiModule.CustomerApiModuleDeps{
			Router:                      router,
			CustomerSubscriptionService: customerSubscriptionService,
			ValidatorX:                  validatorX,
		},
	)

	//return stack(router)
	return &AppReturningValue{
		HttpHandler: stack(router),
		//GrpcServer:  grpcServer,
		Services: AppServicesType{},
	}
}
