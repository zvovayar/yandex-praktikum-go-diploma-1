export RUN_ADDRESS=localhost:7777
export DATABASE_URI=postgres://postgres:qweasd@localhost:5432/yandexdiploma1?sslmode=disable
export ACCRUAL_SYSTEM_ADDRESS=http://localhost:8080

go run ./cmd/gophermart/main.go -v +