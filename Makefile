build:
	go build -o bin/tasks cmd/tasks/main.go

gen:
	rm -rf gen
	buf generate
	
PHONY:
	build