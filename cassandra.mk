DOCKER_COMPOSE := docker compose

cassandra-up:
	$(DOCKER_COMPOSE) up -d cassandra

cassandra-down:
	$(DOCKER_COMPOSE) down cassandra
