run :
	nodemon --exec go run main.go --signal SIGTERM

swag:
	swag init
	# swag init -g internal/product/handler/http/handler.go
