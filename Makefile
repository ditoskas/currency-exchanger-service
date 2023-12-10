BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Generate build
	go build -o ${BIN_DIR}/${SERVER_DIR} ./${SERVER_DIR}
	go build -o ${BIN_DIR}/${CLIENT_DIR} ./${CLIENT_DIR}

protoc: ## Generate Pbs
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto

protoc-go:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE}/proto --go_out=${SERVER_DIR}/pb --go-grpc_opt=module=${PACKAGE}/proto --go-grpc_out=${SERVER_DIR}/pb ${PROTO_DIR}/*.proto

protoc-php: ##Generate PHP files
	protoc -I${PROTO_DIR} --php_out=${CLIENT_DIR}/src --grpc_out=${CLIENT_DIR}/src --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin ${PROTO_DIR}/*.proto

clean: ## Clean generated files
	rm -rf ${PROTO_DIR}/*.pb.go
	rm -rf ${BIN_DIR}

prod: clean protoc build ## Clean and build for production

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}