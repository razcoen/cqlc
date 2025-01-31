DOCKER_COMPOSE=docker compose
GO_RUN=go run

cassandra-up:
	$(DOCKER_COMPOSE) up -d cassandra

cassandra-down:
	$(DOCKER_COMPOSE) down cassandra

cassandra-healthcheck:
	$(GO_RUN) ./scripts/cassandra-healthcheck.go -sleep 1s -timeout 2m
