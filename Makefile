.PHONY: bcjoy

build-production: static
	@go build

static: template/*
	@staticfiles -o files/files.go template/
