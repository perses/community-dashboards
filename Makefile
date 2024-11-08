.PHONY: demo
start-demo:
	@echo "Setting up demo environment"

	@cd ./examples && docker-compose up -d

.PHONY: clean-demo
clean-demo:
	@echo "Cleaning up demo environment"

	@cd ./examples && docker-compose down -v

.PHONY: build-dashboards
build-dashboards:
	@echo "Building dashboards"
	@go run main.go