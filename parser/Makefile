PARSER_CONTAINER := go-parser
PARSER_DIR=./cmd/parser

build-container:
	@sh ./build/build.sh

start-container:
	@echo "running container..."
	@docker run -d --name $(PARSER_CONTAINER) $(PARSER_CONTAINER) sleep 38768
	@echo "container is up!"

container-scrap:
	@echo "Start scrapping..."
	@go run $(PARSER_DIR)/main.go

scrap:
	@docker exec -ti $(PARSER_CONTAINER) make container-scrap

copy:
	@sh ./build/copy.sh

lint: 
	@echo "start linting..."
	@./build/lint.sh
	@echo "linting finished."

