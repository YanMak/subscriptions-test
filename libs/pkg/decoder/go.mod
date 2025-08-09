module project1.v0/pkg/decoder

go 1.24.4

require (
	github.com/google/uuid v1.6.0
	project1.v0/contracts/transport/http v0.0.0
	project1.v0/pkg/validator_x v0.0.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.27.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	project1.v0/contracts/domain v0.0.0 // indirect
)

replace project1.v0/contracts/transport/http => ../../contracts/transport/http

replace project1.v0/pkg/validator_x => ../validator-x

replace project1.v0/contracts/domain => ../../contracts/domain
