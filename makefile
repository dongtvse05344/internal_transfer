access_db:
	docker exec -it transfer_db mysql -uroot -pmy_password

start_db:
	docker-compose up -d

migrate_up:
	migrate -path dal/db/migration -database "mysql://root:my_password@tcp(127.0.0.1:3306)/mysql" -verbose up

migrate_down:
	migrate -path dal/db/migration -database "mysql://root:my_password@tcp(127.0.0.1:3306)/mysql" -verbose down
init_env:
	brew install golang-migrate

sqlc:
	sqlc generate
proto_gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
               --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
               --grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
               proto/*.proto