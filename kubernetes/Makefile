.PHONY: docker
docker:
	@rm webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -o webook web/main.go
	@docker rmi -f rui-cs/webook:v0.0.1
	@docker build -t rui-cs/webook:v0.0.1 .
