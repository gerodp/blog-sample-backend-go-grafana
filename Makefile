
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

starttest:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose -f docker-compose-test.yml build && docker compose -f docker-compose-test.yml up

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

entertestdb:
	docker compose exec -it mariadb bash -c 'mysql -u root -ppassword blogTest'

enterb:
	docker compose exec -it backend bash


createMariaDBPromExporterUser:
	docker compose exec -it mariadb bash -c 'mysql -u root -ppassword -e "CREATE USER '\''exporter1'\''@'\''%'\'' IDENTIFIED BY '\''password'\'' WITH MAX_USER_CONNECTIONS 3;GRANT PROCESS, REPLICATION CLIENT ON *.* TO '\''exporter1'\''@'\''%'\'';GRANT SELECT ON performance_schema.* TO '\''exporter1'\''@'\''%'\'';"'


createdb:
	docker compose exec -it mariadb bash -c 'mysql -u root -ppassword -e "CREATE DATABASE blog;"'



createtestdb:
	docker compose exec -it mariadb bash -c 'mysql -u root -ppassword -e "CREATE DATABASE blogTest;"'

help:
	cat Makefile
