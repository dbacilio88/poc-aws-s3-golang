test:
	go test -v -c ./...


plan:
	workspace/terraform/ terraform init

.PHONY:

