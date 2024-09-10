run:
	go mod tidy
	clear
	go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ muallimah-submodule/protos/*.proto
	

migrate_up:
	migrate -path muallimah-submodule/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/muallimah_db?sslmode=disable -verbose up

migrate_down:
	migrate -path muallimah-submodule/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/muallimah_db?sslmode=disable -verbose down


migrate_force:
	migrate -path muallimah-submodule/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/muallimah_db?sslmode=disable -verbose force 1


migrate_file:
	migrate create -ext sql -dir muallimah-submodule/migrations -seq insert_table

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1
