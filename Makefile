## Configure Go build options
GIN_MODE = "release"
CGO_ENABLED = 0

## Configure DB connection params
DB = "root:password@(localhost:3306)/library"

run:
	 cd app && GIN_MODE=${GIN_MODE} DB=${DB} go run main.go

lint:
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ./...

build:
	# MacOS
	# cd app && GOOS=darwin GOARCH=amd64 go build -o ../bin/main-darwin-amd64 main.go
	# Linux
	cd app && GOOS=linux GOARCH=amd64 go build -o ../bin/main-linux-amd64 main.go
	# Windows
	# cd app && GOOS=windows GOARCH=amd64 go build -o ../../bin/main-windows-amd64 main.go

docker_up:
	docker build -t api:0.1.0 . 
	docker-compose -f docker/docker-compose.api.yml up -d
	docker logs docker_app_1 