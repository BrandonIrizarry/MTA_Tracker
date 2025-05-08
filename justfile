build-all:
	go build -o stopmonitortest cmd/stopmonitortest/main.go

build cmd:
	go build -o {{cmd}} cmd/{{cmd}}/main.go

run cmd: (build cmd)
	./{{cmd}}

view-db cmd:
	sqlite3 cmd/{{cmd}}/{{cmd}}.db

sqlc-gen cmd:
	sqlc -f cmd/{{cmd}}/sqlc.yaml

reset-db cmd:
	rm -rf cmd/{{cmd}}/{{cmd}}.db
	goose -dir cmd/{{cmd}}/sql/schema sqlite3 cmd/{{cmd}}/{{cmd}}.db up
