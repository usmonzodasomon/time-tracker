run:
	go run cmd/time-tracker/main.go

migrate_up:
	goose -dir migrations postgres "postgresql://postgres:qwerty@127.0.0.1:5432/timetracker?sslmode=disable" up

migrate_down:
	goose -dir migrations postgres "postgresql://postgres:qwerty@127.0.0.1:5432/timetracker?sslmode=disable" down