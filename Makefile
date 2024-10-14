build:
	@go build ./bin/main.exe

run:
	@go run ./cmd/main/main.go

test-ast:
	@go test ./ast

test-parser:
	@go test ./parser

test-lexer:
	@go test ./internals/lexer
test-object:
	@go test ./internals/objects
test-evaluator:
	@go test ./evaluator
