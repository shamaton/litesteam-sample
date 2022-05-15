up:
	docker compose up

down:
	docker compose down

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
	curl -OL https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-darwin-amd64.zip
	unzip litestream-v0.3.8-darwin-amd64.zip

PHONY: up gen create-table insert-user select-user