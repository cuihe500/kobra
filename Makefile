APP=kobra

.PHONY: clean
clean:
	go clean
	rm -rf kobra kobra.*

.PHONY: build
build:
	go mod tidy
	go build -o ${APP} app/main.go

.PHONY: run
run:
	go run app/main.go server -c config/dev-config.toml