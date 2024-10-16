MAKEFLAGS += -j4

templ:
	templ generate

css:
	bun run tailwind

generate: templ css

watch-template:
	fswatch -or ./domain/template/**.templ | xargs -n1 -I{} make generate

watch-dev:
	gowatch

dev: watch-dev watch-template

tidy:
	go mod tidy

build:
	docker build -t mohammaddocker/newsletter:latest --platform linux/amd64 .

push: build
	docker push mohammaddocker/newsletter

