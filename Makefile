TIMESTAMP             	:= $(shell /bin/date "+%F %T")

default:
	@echo "no default target"; false

fmt:
	@go fmt ./...
	@go mod tidy

github: fmt
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: default fmt github