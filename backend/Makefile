# Makefile for running the JobSphere backend API

# Variables
SOURCE_FILE = ./cmd/api/main.go

# Default target to display help
.DEFAULT_GOAL := help

# Run the application
.PHONY: run
run:
	@echo "Starting the application..."
	go run $(SOURCE_FILE)

# Help message
.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make run       - Run the application"
	@echo "  make help      - Display this help message"