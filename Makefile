
A !=
B !=

ps:
	docker compose ps

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose build

buildverbose:
	BUILDKIT_PROGRESS=plain COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose build

restart:
	docker compose restart

start:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose up
stop:
	docker compose stop

runb:
	go run backend/main.go

logs:
	docker compose logs ${A}

logsb:
	docker compose logs backend

enterdb:
	docker compose exec -it mariadb bash -c 'mysql -u root -ppassword blog'

enterb:
	docker compose exec -it backend bash

uninstall:
	docker compose down

help:
	cat Makefile
