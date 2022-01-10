ENV 					?= dev
build:
	@echo ">> go build ..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags ${ENV} --ldflags -w -o ../fiber-demo