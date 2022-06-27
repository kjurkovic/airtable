create_image: swagger
	@echo "\033[35mCreating auth docker image\033[0m"
	docker build --tag collabtable-auth .

verify_swagger:	
	@echo "\033[35mVerifying Go Swagger installation\033[0m"
	which swagger || GO111MODULE=on go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: verify_swagger
	@echo "\033[35mGenerating auth swagger docs\033[0m"
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models
