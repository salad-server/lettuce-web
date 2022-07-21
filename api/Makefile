# Build
build:
	go build -o mini-api ./cmd/

# Build with upx (compress output)
build-prod:
	go build -o mini-api ./cmd/
	upx mini-api

# Run without building (development)
serve:
	go run cmd/*.go
