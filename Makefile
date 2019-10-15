TIMESTAMP             	:= $(shell /bin/date "+%F %T")

default:
	@echo "no default target"; false

fmt:
	@go fmt ./...
	@go mod tidy

protoc:
	protoc -I=$(CURDIR)/proto/ --go_out=$(CURDIR) $(CURDIR)/proto/snowflake.proto

github: protoc fmt
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: default fmt protoc github
