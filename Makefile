.PHONY: run build
run:
	docker-compose --env-file ./deployments/.env up

build:
	docker-compose --env-file ./deployments/.env build