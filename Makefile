.PHONY: unit int

unit:
	@go test ./... -v -run Test[^Integration]
int:
	@go test ./... -v -run Integration

name=''
test:
	@go test ./... -v -run $(name)