run :
	nodemon --exec go run main.go --signal SIGTERM

swag:
	swag fmt
	swag init -g main.go internal/product/handler/http/handler.go
	# swag init -g internal/product/handler/http/handler.go
