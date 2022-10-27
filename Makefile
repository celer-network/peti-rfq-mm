install:
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify
	go build -o $(HOME)/go/bin/peti-rfq-mm ./example-mm/main/main.go

.PHONY: proto
proto:
	protoc --proto_path=./sdk/proto --go_out ./sdk --go_opt=module=github.com/celer-network/peti-rfq-mm/sdk \
	--go-grpc_out=./sdk --go-grpc_opt=require_unimplemented_servers=false,module=github.com/celer-network/peti-rfq-mm/sdk \
	--grpc-gateway_out ./sdk --grpc-gateway_opt=module=github.com/celer-network/peti-rfq-mm/sdk \
	--openapiv2_out ./sdk/openapi \
	./sdk/proto/service/*/*.proto ./sdk/proto/common/*.proto