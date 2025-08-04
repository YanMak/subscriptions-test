module project1.v0/pkg/db

go 1.24.4

replace project1.v0/configs => ../../configs

replace project1.v0/mappers/conv => ../../mappers/conv

require project1.v0/mappers/conv v0.0.0 // indirect

require (
	github.com/lib/pq v1.10.9
	project1.v0/configs v0.0.0
)
