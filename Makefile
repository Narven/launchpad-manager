include .env
export

.DEFAULT_GOAL := start

start: server

server:
	@echo "\n\t🤖️ Run Server\n"
	go run src/main.go