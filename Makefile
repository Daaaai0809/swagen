run_generate_message_test:
	@echo "Running generate message test"
	@sh scripts/message_generate_test.sh

run_generate_path_test:
	@echo "Running generate path test"
	@sh scripts/path_generate_test.sh

run_test: run_generate_message_test run_generate_path_test
	@echo "Running all tests"

build:
	@echo "Building project"
	@go build -o swagen
