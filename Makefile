.phony: dump-schema
dump-schema:
	psqldef -U veerush -Wpassword -h localhost -p 5432 veerush --export >sqlc/schema.sql

.phony: gen-sqlc
gen-sqlc:
	docker pull sqlc/sqlc:1.24.0
	docker run --rm -v .:/src -w /src sqlc/sqlc generate

.phony: clean-gen-sqlc
clean-gen-sqlc:
	rm -rf internal/database/generated