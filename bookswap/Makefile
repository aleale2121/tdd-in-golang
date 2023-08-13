up:
	docker-compose config
	docker-compose up
migrate_init:
	migrate create -ext sql -dir db/migrations -seq create_users_table
	migrate create -ext sql -dir db/migrations -seq create_books_table