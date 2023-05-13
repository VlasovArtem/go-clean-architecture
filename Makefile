runTests: ##Run Test
	go test ./...

generateMocks: ##Generate Mocks
	go generate ./...

build: ##Build
	go build ./...

buildImage: ##Build Image
	docker build -t avlasov/clean-architecture:v0.0.1 -f ./build/Dockerfile .

analyze: ##Analyze
	go vet ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'