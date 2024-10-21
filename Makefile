include .env

create_migration:
	migrate create -ext sql -dir ./migrations -seq $(name)

migrate_up:
	migrate -dir ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB}" -verbose up

migrate_down:
	migrate -dir ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB}" -verbose down

migrate_down_to:
	migrate -dir ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB}" -verbose down $(version)