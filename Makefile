run:
	godotenv -f ./local.env go run main.go

load-test:
	k6 run ./load-tests/add-free.js -q
	k6 run ./load-tests/add-with-redis-lock.js -q