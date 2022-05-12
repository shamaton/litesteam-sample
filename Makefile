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

PHONY: up gen create-table insert-user select-user