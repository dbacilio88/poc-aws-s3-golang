test:
	go test -v -c ./...


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

.PHONY:

