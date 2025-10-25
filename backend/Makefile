build:
	go build -o bin/chatear-api ./cmd/api
	go build -o bin/notification_worker ./cmd/worker/notification_worker.go
	go build -o bin/user_delete_worker ./cmd/worker/user_delete_worker.go
	go build -o bin/user_hard_delete_worker ./cmd/worker/user_hard_delete_worker.go
	go build -o bin/user_permanent_deletion_scheduler_worker ./cmd/worker/user_permanent_deletion_scheduler_worker.go
	go build -o bin/user_registered_worker ./cmd/worker/user_registered_worker.go

run-api:
	go run ./cmd/api

run-worker-notification:
	go run ./cmd/worker/notification_worker.go

run-worker-user-delete:
	go run ./cmd/worker/user_delete_worker.go

run-worker-user-hard-delete:
	go run ./cmd/worker/user_hard_delete_worker.go

run-worker-user-permanent-deletion-scheduler:
	go run ./cmd/worker/user_permanent_deletion_scheduler_worker.go

run-worker-user-registered:
	go run ./cmd/worker/user_registered_worker.go

test:
	go test ./... -v

lint:
	go vet ./...

# Placeholder for migration commands (assuming golang-migrate or atlas)
migrate/up:
	@echo "Migration command not yet implemented. Please refer to documentation for migration tools (golang-migrate/atlas)."

clean:
	rm -rf bin

.PHONY: build run-api run-worker-notification run-worker-user-delete run-worker-user-hard-delete run-worker-user-permanent-deletion-scheduler run-worker-user-registered test lint migrate/up clean
