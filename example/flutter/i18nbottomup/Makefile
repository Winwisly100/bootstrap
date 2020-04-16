GO_OS 		:= $(shell go env GOOS)
GO_ARCH		:= $(shell go env GOARCH)

GIT_VERSION	:= $(shell git describe --tags)

current_dir = $(shell pwd)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

SAMPLE_NAME = maintemplate
SAMPLE_FSPATH = $(PWD)/$(SAMPLE_NAME)

# Absolute File path to Server
SERVER_FSPATH=$(current_dir)/../server

I18N_TEMPLATE_DIR = $(current_dir)/i18n/
I18N_TEMPLATE_NAME = "intl.arb"
I18N_TEMPLATE_PATH = $(I18N_TEMPLATE_DIR)/$(I18N_TEMPLATE_NAME)
I18N_PREFIX_OUT_FILES = "stock"

## Show help
help:
	@echo 'Provides a better way to handle i18n in Flutter.'
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
## print
print: 
	@echo
	@echo GO_OS: $(GO_OS)
	@echo GO_ARCH: $(GO_ARCH)
	@echo GIT_VERSION: $(GIT_VERSION)
	@echo LIB_FSPATH: $(LIB_FSPATH)
	@echo

	@echo SAMPLE
	@echo SAMPLE_FSPATH: $(SAMPLE_FSPATH)
	@echo

	@echo I18N
	@echo I18N_TEMPLATE_PATH: $(I18N_TEMPLATE_PATH)
	@echo I18N_PREFIX_OUT_FILES: $(I18N_PREFIX_OUT_FILES)
	@echo

	@echo SERVER:
	@echo SERVER_FSPATH: $(SERVER_FSPATH)
	@echo

## flu-config
flu-config: 
	flutter channel beta
	flutter upgrade --force

## flu-web-run
flu-web-run: 
	flutter config --enable-web
	cd $(SAMPLE_FSPATH) && flutter run -d chrome

