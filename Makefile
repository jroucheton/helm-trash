VERSION := "1.0.0"
DIST := $(CURDIR)/_dist
LDFLAGS := "-X main.version=${VERSION}"
TAR_LINUX := "helm-trash-linux-amd64.tar.gz"
TAR_WINDOWS := "helm-trash-windows-amd64.tar.gz"
BINARY_LINUX := "helm-trash"
BINARY_WINDOWS := "helm-trash.exe"

.PHONY: dist

dist: dist_linux dist_windows

dist_linux:
	mkdir -p $(DIST)
	GOOS=linux GOARCH=amd64 go get -t -v ./...
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX) -ldflags $(LDFLAGS) main.go
	sed -e "s/:VERSION:/$(VERSION)/g" -e "s/:BINARY:/$(BINARY_LINUX)/g" plugin.yaml.template > plugin.yaml
	tar -czvf $(DIST)/$(TAR_LINUX) $(BINARY_LINUX) README.md LICENSE plugin.yaml

.PHONY: dist_windows
dist_windows:
	GOOS=windows GOARCH=amd64 go get -t -v ./...
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_WINDOWS) -ldflags $(LDFLAGS) main.go
	sed -e "s/:VERSION:/$(VERSION)/g" -e "s/:BINARY:/$(BINARY_WINDOWS)/g" plugin.yaml.template > plugin.yaml
	tar -czvf $(DIST)/${TAR_WINDOWS} $(BINARY_WINDOWS) README.md LICENSE plugin.yaml
