.PHONY: run_generate_message_test
run_generate_message_test:
	@echo "Running generate message test"
	@sh scripts/message_generate_test.sh

.PHONY: run_generate_path_test
run_generate_path_test:
	@echo "Running generate path test"
	@sh scripts/path_generate_test.sh

.PHONY: run_test
run_test: run_generate_message_test run_generate_path_test
	@echo "Running all tests"

.PHONY: build
build: run_test
	@echo "Building project"
	@go build -o swagen
