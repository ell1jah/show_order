.PHONY: run build publish
run:
	docker-compose --env-file ./deployments/.env up

build:
	docker-compose --env-file ./deployments/.env build

publish:
	go run tools/nats_publisher/main.go