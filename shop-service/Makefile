run:
	go mod tidy
	clear
	go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ muallimah-submodule/protos/*.proto

migrate_up:
	migrate -path migrations -database postgres://postgres:root@localhost:5432/lingua_progress?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:root@localhost:5432/lingua_progress?sslmode=disable -verbose down

migrate_force:th migrations -database postgres://postgres:root@localhost:5432/lingua_progress?sslmode=disable -verbose force 1

migrate_file:
	migrate -pa
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1