up:
	docker compose up

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

PHONY: up gen create-table insert-user select-user