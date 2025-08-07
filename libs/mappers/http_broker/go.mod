module project1.v0/mappers/http_broker

go 1.24.4

replace project1.v0/mappers/conv => ../conv

replace project1.v0/contracts/transport/broker => ../../contracts/transport/broker

replace project1.v0/contracts/transport/http => ../../contracts/transport/http

replace project1.v0/contracts/domain => ../../contracts/domain

require (
	project1.v0/contracts/transport/broker v0.0.0
	//project1.v0/mappers/conv v0.0.0
	project1.v0/contracts/transport/http v0.0.0
)

require (
	github.com/google/uuid v1.6.0 // indirect
	project1.v0/contracts/domain v0.0.0 // indirect
)
