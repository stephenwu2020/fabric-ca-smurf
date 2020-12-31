.PHONY: start-ca-server
start-ca-server:
	docker-compose -f caserver/docker-compose.yaml up -d