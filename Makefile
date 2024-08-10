PORT := ":8080"

binary:
	@go build ./cmd/account-server/*

run: binary
	@./account-server --port $(PORT)