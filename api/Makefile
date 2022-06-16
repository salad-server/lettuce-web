build:
	go build -o mini-api ./cmd/

build-prod:
	go build -o mini-api ./cmd/
	upx mini-api

serve:
	go run cmd/*.go

stop:
	killall mini-api
