PARSER_CONTAINER := go-parser
PARSER_DIR=./cmd/parser
PARSER_BUILD_FILE=parser

SCRAP_URL_WIKI=https://uk.wikiquote.org/wiki/%D0%A8%D0%B5%D0%B2%D1%87%D0%B5%D0%BD%D0%BA%D0%BE_%D0%A2%D0%B0%D1%80%D0%B0%D1%81_%D0%93%D1%80%D0%B8%D0%B3%D0%BE%D1%80%D0%BE%D0%B2%D0%B8%D1%87

build-container:
	@sh ./build/build.sh

start-container:
	@echo "running container..."
	@docker run -d --name $(PARSER_CONTAINER) $(PARSER_CONTAINER) sleep 38768
	@echo "container is up!"

build-executable:
	@echo "building parser..."
	@docker exec -i $(PARSER_CONTAINER) go build -o $(PARSER_BUILD_FILE) $(PARSER_DIR)/main.go
	@echo "parser built"

scrap-wiki-quotes:
	@echo "start scrapping from $(SCRAP_URL_WIKI)..."
	@docker exec -i $(PARSER_CONTAINER) ./parser parse --url $(SCRAP_URL_WIKI) --from wiki-quotes

copy:
	@sh ./build/copy.sh

copy-windows:
	@sh ./build/copy-win.sh

lint: 
	@echo "start linting..."
	@./build/lint.sh
	@echo "linting finished."
