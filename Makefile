PORT := ":8080"

build:binary
	@docker build
binary:
	@go build ./cmd/app/*

run: binary
	@./app-server --port $(PORT)

run-docker:
	@docker build --tag superm4n/api-server .
	@docker run -dp 8080:8080 --name my-api-server superm4n/api-server
	@docker container rm -f my-api-server || true
	@docker start my-api-server -a
#you can exec into the container using
#	@docker exec -it my-account-server sh
