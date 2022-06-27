create_image: swagger
	@echo "\033[32mCreating workspace docker image\033[0m"
	docker build --tag collabtable-workspace .

verify_install:	
	@echo "\033[32mVerifying Go Swagger installation\033[0m"
	which swagger || GO111MODULE=on go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: verify_install
	@echo "\033[32mGenerating workspace swagger docs\033[0m"
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models
