# Run Go unit tests
test-go:
	go test ./...

# Run Python unit tests
test-py:
	@for dir in medium/*/; do \
		if [ -f "$$dir/test_main.py" ]; then \
			.venv/bin/python -m pytest "$$dir/test_main.py" -v; \
		fi \
	done

# Run all unit tests (Go first, then Python)
test: test-go test-py
