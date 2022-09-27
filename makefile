dev:
		CompileDaemon -command="./go-ch-test"
test:
		go test -v --cover ./...