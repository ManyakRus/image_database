SERVICENAME=image_database
SERVICEURL=github.com/ManyakRus/$(SERVICENAME)

FILEMAIN=./internal/main.go
FILEAPP=./bin/$(SERVICENAME)

NEW_REPO=github.com/ManyakRus/image_database


run:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	./bin/$(SERVICENAME)
mod:
	clear
	go get -u ./...
	go mod tidy -compat=1.22
	go mod vendor
	go fmt ./...
build:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	cd ./cmd && \
	./VersionToFile.py
	cp $(FILEAPP) $(GOPATH)/bin

lint:
	clear
	go fmt ./...
	golangci-lint run ./internal/...
	golangci-lint run ./pkg/...
	gocyclo -over 10 ./internal
	gocyclo -over 10 ./pkg
	gocritic check ./internal/...
	gocritic check ./pkg/...
	staticcheck ./internal/...
	staticcheck ./pkg/...
run.test:
	clear
	go fmt ./...
	go test -coverprofile cover.out ./internal/v0/app/...
	go tool cover -func=cover.out
newrepo:
	sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g' go.mod
	find -name *.go -not -path "*/vendor/*"|xargs sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g'
graph:
	clear
	image_packages ./ docs/packages.graphml
conn:
	clear
	image_connections ./internal docs/connections.graphml $(SERVICENAME)
lines:
	clear
	go_lines_count ./ ./docs/lines_count.txt 10
licenses:
	golicense -out-xlsx=./docs/licenses.xlsx $(FILEAPP)
gocyclo:
	golangci-lint run ./... --disable-all -E gocyclo -v
