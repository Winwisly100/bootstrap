

# ALL VARS get overidden in your own "env.mk"

SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# Grab necessary submodules, in case the repo was cloned without --recursive
#$(SELF_DIR)/.ci/common.mk:
#	git submodule update --init --recursive

## TODO
# Make it work with example.
# What order to run the flutter i18n in ?
# Make file does not actually run the example and do the required translation things on it.
# This was meant to be sorted ages ago.

# Incldue your env.mk
include $(SELF_DIR)/env.mk


help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


print: ## print

	@echo
	@echo SELF_DIR: $($(SELF_DIR))
	
	@echo

	@echo
	@echo LIB_BIN_NAME: $(LIB_BIN_NAME)
	@echo LIB_BIN_FSPATH: $(LIB_BIN_FSPATH)
	@echo LIB_CONFIG_FSPATH: $(LIB_CONFIG_FSPATH)
	@echo SAMPLE_FSPATH: $(SAMPLE_FSPATH)
	
	@echo

	@echo
	@echo GOOGLE_SHEET_ID: $(GOOGLE_SHEET_ID)
	@echo GOOGLE_SHEET_URL: $(GOOGLE_SHEET_URL)
	@echo GOOGLE_SHEET_CSV: $(GOOGLE_SHEET_CSV)
	@echo

code:
	code $(LIB_FSPATH)

dep:
	# none

test:
	go test -v ./...

	
### cloud

cloud-auth:
	gcloud auth login
	gcloud config set project $(CLOUD_PROJECT_ID)

	gcloud config set run/region europe-west1
cloud-open:
	# run
	open https://console.cloud.google.com/run?project=$(CLOUD_PROJECT_ID)
	# container
	open https://console.cloud.google.com/cloud-build/builds?project=$(CLOUD_PROJECT_ID)
	# store
	open https://console.cloud.google.com/storage/browser?project=$(CLOUD_PROJECT_ID)
cloud-build:
	echo $(CLOUD_PROJECT_ID)
	gcloud builds submit --tag gcr.io/$(CLOUD_PROJECT_ID)/identicon-generator
cloud-deploy:
	gcloud beta run deploy --image gcr.io/$(CLOUD_PROJECT_ID)/identicon-generator --platform managed



gsheet-open: ## gsheet-open
	open $(GOOGLE_SHEET_URL)
	open $(GOOGLE_SHEET_CSV)

### build
build: ## build ****
	mkdir -p $(LIB_CONFIG_FSPATH)
	go build -o $(LIB_BIN_FSPATH) .
	cp ./config/* $(LIB_CONFIG_FSPATH)

build-clean:
	# go stuff
	go clean -cache
	go clean -modcache
	go clean -testcache

	# binary in the GO BIN path
	rm $(LIB_BIN_FSPATH)

run-datadump: ## run-datadump
	#rm -rf ./outputs
	$(LIB_BIN_FSPATH) gsheet -o datadump

run-langdump: ## run-langdump
	#rm -rf ./outputs
	$(LIB_BIN_FSPATH) gsheet -o lang
	
run-hugo: ## run-langdump
	#rm -rf ./outputs
	$(LIB_BIN_FSPATH) gsheet -o hugo

# i18n tool
# 1. make build ==> build i18n tool
# 2. make i18n-full-json-flutter ==> generate translated json files
# 3. make i18n-arb-flutter ==> convert json to arb

i18n-full-json-flutter:
	# generate full json files
	$(LIB_BIN_FSPATH) flutter --dir examples/i18n_example/lib/i18n/generated/ --template examples/i18n_example/lib/i18n/generated/intl_messages.arb --prefix mockText --languages en,fr,es,ur -f

i18n-arb-flutter:
	$(LIB_BIN_FSPATH) flutter --dir examples/i18n_example/lib/i18n/generated/
