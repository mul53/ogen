GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
OGEN_VERSION=1.0.0

ifeq ($(OS),Windows_NT)
    OS := Windows
else
    OS := $(shell uname)  # same as "uname -s"
endif

LOWECASE_OS = $(shell echo $(OS) | tr A-Z a-z)
FOLDER_NAME= $(BINARY_NAME)-$(OGEN_VERSION)

ifeq ($(LOWECASE_OS),darwin)
    BINARY_NAME := ogen
else ifeq ($(LOWECASE_OS),linux)
    BINARY_NAME := ogen
else ifeq ($(LOWECASE_OS),windows)
    BINARY_NAME := ogen.exe
endif

run: build
	@echo Running $(BINARY_NAME)
	./$(BINARY_NAME)

pack-windows-amd64: build-windows-amd64
	mkdir $(FOLDER_NAME)
	mv $(BINARY_NAME) ./$(FOLDER_NAME)
	zip -r ogen-$(OGEN_VERSION)-windows-amd64.zip ./$(FOLDER_NAME)
	rm -r ./$(FOLDER_NAME)

build-windows-amd64:
	@echo Building $(BINARY_NAME) for windows amd64
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)	

pack-linux-arm64: build-linux-arm64
	mkdir $(FOLDER_NAME)
	mv $(BINARY_NAME) ./$(FOLDER_NAME)
	tar -czvf ogen-$(OGEN_VERSION)-linux-arm64.tar.gz ./$(FOLDER_NAME)
	rm -r ./$(FOLDER_NAME)

build-linux-arm64:
	@echo Building $(BINARY_NAME) for linux arm64
	env CGO_ENABLED=1 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BINARY_NAME)

pack-linux-amd64: build-linux-amd64
	mkdir $(FOLDER_NAME)
	mv $(BINARY_NAME) ./$(FOLDER_NAME)
	tar -czvf ogen-$(OGEN_VERSION)-linux-amd64.tar.gz ./$(FOLDER_NAME)
	rm -r ./$(FOLDER_NAME)

build-linux-amd64:
	@echo Building $(BINARY_NAME) for linux amd64
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)

pack-darwin: build-darwin
	mkdir $(FOLDER_NAME)
	mv $(BINARY_NAME) ./$(FOLDER_NAME)
	tar -czvf ogen-$(OGEN_VERSION)-darwin.tar.gz ./$(FOLDER_NAME)
	rm -r ./$(FOLDER_NAME)

build-darwin:
	@echo Building $(BINARY_NAME) for darwin
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)
	
release: clean pack-darwin pack-linux-amd64 pack-linux-arm64 pack-windows-amd64
	mkdir ./release
	mv ogen-$(OGEN_VERSION)-* ./release

pack: build
	mkdir $(FOLDER_NAME)
	mv $(BINARY_NAME) ./$(FOLDER_NAME)
	tar -czvf ogen-$(OGEN_VERSION)-$(LOWECASE_OS).tar.gz ./$(FOLDER_NAME)
	rm -r ./$(FOLDER_NAME)

build: 
	@echo Building $(BINARY_NAME) for $(OS)
	$(GOBUILD) -o $(BINARY_NAME)

clean:
	@echo Cleaning...
	$(GOCLEAN) ./
	rm -rf ./$(BINARY_NAME)
	rm -rf ./builds
	rm -rf ogen-darwin*
	rm -rf ogen-windows*
	rm -rf ogen-linux*
	rm -rf *.tar.gz
	rm -rf *.zip
	rm -rf release/
	rm -rf ogen-protopy*




