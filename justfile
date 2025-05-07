build-all:
	go build -o stopmonitortest cmd/stopmonitortest/main.go

build cmd:
	go build -o {{cmd}} cmd/{{cmd}}/main.go

run cmd: (build cmd)
	./{{cmd}}
