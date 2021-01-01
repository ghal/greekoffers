help:
	@echo "Please use 'make <target>' where <target> is one of the following:"
	@echo "  run             to run the app."
	@echo "  lint            to perform linting."

run:
	docker-compose -f ./deployments/docker-compose.yaml up greekoffers

stop:
	docker-compose down

lint:
	golint -set_exit_status=1 `go list ./...`