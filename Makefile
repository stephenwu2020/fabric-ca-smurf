.PHONY: start-ca-server
start-ca-server:
	docker-compose -f caserver/docker-compose.yaml up -d

.PHONY: stop-ca-server
stop-ca-server:
	docker-compose -f caserver/docker-compose.yaml down