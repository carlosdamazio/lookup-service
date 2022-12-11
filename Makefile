build:
	docker build -t lookupsvc:latest .

localdev-start:
	docker compose up -d --build

test:
	go test ./...