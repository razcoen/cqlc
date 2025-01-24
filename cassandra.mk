DOCKER_COMPOSE := docker compose

cassandra-up:
	$(DOCKER_COMPOSE) up cassandra

cassandra-down:
	$(DOCKER_COMPOSE) down cassandra
