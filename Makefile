create_migration:
	migrate create -ext=sql -dir=sql/migrations -seq $(name)

migrateup:
		migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
		migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: create_migration migrateup migratedown