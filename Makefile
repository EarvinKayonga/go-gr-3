build:
	go build -o bin/tasks cmd/tasks/main.go

gen:
	buf generate
	
PHONY:
	build