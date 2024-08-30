

DBURL="postgres://postgres:root@localhost:5432/delivery_product_db?sslmode=disable"
exp:
	export DBURL="postgres://postgres:root@localhost:5432/delivery_product_db?sslmode=disable"

mig-up:
	migrate -path muallimah-submodule/migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path muallimah-submodule/migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir muallimah-submodule/migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir muallimah-submodule/migrations -seq insert_table

CURRENT_DIR=$(shell pwd)

proto-gen:
	./internal/pkg/scripts/gen-proto.sh $(CURRENT_DIR)
	
swag-gen:
	~/go/bin/swag init -g ./internal/http/api.go -o docs force 1	