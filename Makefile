.PHONY: swag_init

swag_init:
	swag init -g ./apps/go/api-gateway/cmd/main.go -o ./libs/contracts/transport/http/swagger/docs