TEST_URL:=http://localhost:3000
up:
	docker-compose up --build -d
down:
	docker-compose down

migrate_init:
	migrate create -ext sql -dir db/migrations -seq create_users_table
	migrate create -ext sql -dir db/migrations -seq create_books_table

test_all:
	go test -run ./...

godog:
	godog run

test_e2e:
	curl --location --request GET $(TEST_URL)
