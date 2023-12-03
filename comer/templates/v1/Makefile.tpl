BINARY_NAME=comer-example
build:
	go build -o ./${BINARY_NAME} -v
	./${BINARY_NAME} server -c configs/settings-local.yml
test:
	cd tests && go test -short `go list ./..`