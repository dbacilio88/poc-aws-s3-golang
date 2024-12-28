test:
	go test ./... -cover -v

report:
	go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml

init:
	terraform -chdir=workspace/terraform init

plan:
	terraform -chdir=workspace/terraform plan

apply:
	terraform -chdir=workspace/terraform apply

destroy:
	terraform -chdir=workspace/terraform destroy

update:
	go get -d -v -u all

mock:
	mockgen -source=pkg/env/properties.go -destination=test/mock/pkg/env/properties.go
	mockgen -source=pkg/logs/logger.go -destination=test/mock/pkg/logs/logger.go


.PHONY: test

