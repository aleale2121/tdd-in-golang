BOOKSWAP_BASE_URL:=http://localhost:3000/
BOOKSWAP_PORT=3000
BOOKSWAP_DB_URL:=postgres://root:root@localhost:5432/books?sslmode=disable
up_build:
	docker-compose up --build -d
up:
	docker-compose up -d
down:
	docker-compose down

migrate_init:
	migrate create -ext sql -dir db/migrations -seq create_users_table
	migrate create -ext sql -dir db/migrations -seq create_books_table

install_graphvz:
	sudo apt install graphviz
install_pact:
	go get github.com/pact-foundation/pact-go

mock:
	mockery --dir . --output ./mocks --all
test_all:
	BOOKSWAP_DB_URL=$(BOOKSWAP_DB_URL) bash -c 'go test ./... -v'

test_performance:
	BOOKSWAP_BASE_URL=$(BOOKSWAP_BASE_URL) BOOKSWAP_PORT=$(BOOKSWAP_PORT) bash -c 'go test -bench BenchmarkGetIndex -cpuprofile cpu-books.out ./performance'
	go tool pprof performance.test cpu-books.out
	# type top5
test_performance:
	BOOKSWAP_BASE_URL=$(BOOKSWAP_BASE_URL) BOOKSWAP_PORT=$(BOOKSWAP_PORT) bash -c 'go test -bench BenchmarkGetIndex -cpuprofile cpu-books.out ./performance'
	go tool pprof performance.test cpu-books.out
	# type top5
test_performance_server:
	curl --output book-swap-app.pprof "http://localhost:$(BOOKSWAP_PORT)/debug/pprof/profile?seconds=10"
	# go tool pprof book-swap-app
godog:
	godog run

test_e2e:
	curl --location --request GET $(TEST_URL)
