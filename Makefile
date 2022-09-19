dep:
	go mod download
	
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

format:
	gofmt -s -w .

lint:
	staticcheck ./...

docker_build: build
	sudo docker build .

run: build docker_build
	sudo docker-compose up 

stop:
	sudo docker-compose down 
