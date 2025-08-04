module project1.v0/api_gateway

go 1.24.4

replace (
	project1.v0/configs => ../../../libs/configs
	project1.v0/contracts/domain => ../../../libs/contracts/domain
	project1.v0/contracts/transport/broker => ../../../libs/contracts/transport/broker
	project1.v0/contracts/transport/http => ../../../libs/contracts/transport/http
	project1.v0/mappers/conv => ../../../libs/mappers/conv
	project1.v0/mappers/dto_db => ../../../libs/mappers/dto_db
	project1.v0/mappers/http_broker => ../../../libs/mappers/http_broker
	project1.v0/pkg/db => ../../../libs/pkg/db
	project1.v0/pkg/http => ../../../libs/pkg/http
	project1.v0/pkg/meta => ../../../libs/pkg/meta
	project1.v0/pkg/validator_x => ../../../libs/pkg/validator-x

)

require (
	project1.v0/configs v0.0.0
	//	project1.v0/contracts/transport/http v0.0.0

	project1.v0/contracts/domain v0.0.0 // indirect
	project1.v0/contracts/transport/broker v0.0.0
	project1.v0/contracts/transport/http v0.0.0
	project1.v0/mappers/conv v0.0.0
	project1.v0/mappers/dto_db v0.0.0
	project1.v0/mappers/http_broker v0.0.0
	project1.v0/pkg/db v0.0.0
	project1.v0/pkg/http v0.0.0
	project1.v0/pkg/meta v0.0.0 // indirect
	project1.v0/pkg/validator_x v0.0.0

)

require github.com/google/uuid v1.6.0

require (
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.27.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)
