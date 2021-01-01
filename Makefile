.PHONY: caup
caup:
	docker-compose -f caserver/docker-compose.yaml up -d

.PHONY: cadown
cadown:
	docker-compose -f caserver/docker-compose.yaml down