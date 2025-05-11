build-all:
	go build -o stopmonitortest cmd/stopmonitortest/main.go

build cmd:
	go build -C cmd/{{cmd}}/ -o {{cmd}}.out

run cmd *FLAGS: (build cmd)
	cmd/{{cmd}}/{{cmd}}.out {{FLAGS}}

view-db cmd:
	sqlite3 cmd/{{cmd}}/{{cmd}}.db

sqlc-gen cmd:
	sqlc -f cmd/{{cmd}}/sqlc.yaml generate

reset-db cmd: && (sqlc-gen cmd)
	rm -rf cmd/{{cmd}}/{{cmd}}.db
	goose -dir cmd/{{cmd}}/sql/schema sqlite3 cmd/{{cmd}}/{{cmd}}.db up
