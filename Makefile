PORT := 8090

binary:
	@go build ./cmd/account-server/*

run-binary: binary
	@./account-server --port $(PORT)