APP=kobra

.PHONY: clean
clean:
	go clean
	rm -rf kobra kobra.*

.PHONY: build
build:
	go build -o ${APP} app/main.go

.PHONY: run
run:
	go run app/main.go server -c dev-config.toml