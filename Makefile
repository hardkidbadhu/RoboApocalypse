## Format source code
format:
	go fmt ./...

install_deps:
	go get ./... && go mod vendor

generate_docs:
	swag init

build:
	go build -o robot_apocalypse

build_and_run:
	go build -o robot_apocalypse && ./robot_apocalypse
run:
	go run main.go

db_up:
	docker-compose up -d

db_down:
	docker-compose down
