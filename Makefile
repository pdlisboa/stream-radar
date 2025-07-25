swagger:
	go install github.com/swaggo/swag/cmd/swag@latest && \
 	swag init --dir . --output ./api/docs  --generalInfo ./cmd/api/main.go