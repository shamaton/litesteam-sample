up:
	bash -c "trap 'docker compose down; exit 0' EXIT; docker compose up"

build:
	docker compose build

buildup: build up

gen:
	go generate ./ent/generator/

setup:
	rm -rf ./tmp/certs
	mkdir -p "./tmp/certs"
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./tmp/certs/private.key -out ./tmp/certs/public.crt

PHONY: up build buildup gen setup