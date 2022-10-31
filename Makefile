.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml

.PHONY: formate
formate:
	gofumpt -l -w .

.PHONY: swagger
swagger:
	swag init -g ./internal/adapters/http/api.go
	swagger generate cli -f docs/swagger.json -A tasks_app