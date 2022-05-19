up:
	bash -c "trap 'docker compose down; exit 0' EXIT; docker compose up --remove-orphans"

down:
	docker compose down

downup:	down up

gen:
	go generate ./ent/generator/

create-table:
	CMD=create go run .

insert-user:
	CMD=insert go run .

select-user:
	CMD=select go run .

certkey:
	mkdir -p "./tmp/certs"
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./tmp/certs/private.key -out ./tmp/certs/public.crt

setup:
	rm -rf tmp/sqlite tmp/minio
	go run ./cmd/setup
	docker compose up -d minio minio-client
	docker compose -f docker-compose.yml run --no-deps litestream-replicate
	docker compose down

PHONY: up gen create-table insert-user select-user