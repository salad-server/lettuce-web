# Build
build:
	echo "building mini-api..."
	go build -o mini-api ./cmd/

# Build with upx (compress output)
build-prod:
	echo "building mini-api... (with upx)"
	go build -o mini-api ./cmd/
	upx mini-api

# Run without building (development)
serve:
	echo "serving api..."
	go run cmd/*.go
