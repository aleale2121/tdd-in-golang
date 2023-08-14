TEST_URL:=http://localhost:3000
up:
	docker-compose config
	docker-compose up -d
migrate_init:
	migrate create -ext sql -dir db/migrations -seq create_users_table
	migrate create -ext sql -dir db/migrations -seq create_books_table

godog:
	godog run
test_e2e:
	curl --location --request GET $(TEST_URL)
