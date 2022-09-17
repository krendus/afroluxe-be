run:
	nodemon --exec go run main.go --signal SIGTERM
test:
	go test