PARSER_CONTAINER := go-parser
PARSER_DIR=./cmd/parser

build-container:
	@sh ./build/build_docker.sh

container-scrap:
	@echo "Start scrapping..."
	@go run $(PARSER_DIR)/main.go

copy:
	@sh ./build/copy.sh

start-container:
	@echo "running container..."
	@docker run -d --name $(PARSER_CONTAINER) $(PARSER_CONTAINER) sleep 38768
	@echo "container is up!"

scrap:
	@docker exec -ti $(PARSER_CONTAINER) make container-scrap

lint: 
	@echo "start linting..."
	@./build/lint.sh
	@echo "linting finished."

