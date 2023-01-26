BINARY_NAME=ragbag
BINARY_VERSION=next
SPEC_FILE_LOCATION=./app/docs/swagger.json

build: clean generate
	cd frontend && yarn install && yarn build && cd ..
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-darwin-amd64 app/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-linux-amd64 app/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(BINARY_VERSION)" -o build/${BINARY_NAME}-windows-amd64.exe app/main.go
#	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.version=$(VERSION)" -o build/${BINARY_NAME}-linux-arm main.go

generate: generate-swagger-docs generate-swagger-client

generate-swagger-docs:
	echo generating swagger docs
	swag init -g app/main.go --output app/docs

generate-swagger-client:
	echo generating swagger client
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i "/local/${SPEC_FILE_LOCATION}" \
		-g typescript-fetch \
		--additional-properties=typescriptThreePlus=true \
		-o "/local/frontend/src/gen"
