#!/bin/bash

# Test script for freelancer features
# This script provides multiple ways to test your new freelancer functionality

set -e

echo "🧪 Testing Freelancer Features"
echo "=============================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    print_error "Please run this script from the project root directory"
    exit 1
fi

# Install test dependencies
print_status "Installing test dependencies..."
go mod tidy
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/mock

# Run unit tests
print_status "Running unit tests for freelancer service..."
if go test ./internal/service/freelancer/ -v; then
    print_success "Unit tests passed!"
else
    print_error "Unit tests failed!"
    exit 1
fi

# Run all tests
print_status "Running all project tests..."
if make test; then
    print_success "All tests passed!"
else
    print_error "Some tests failed!"
    exit 1
fi

# Check if Docker is available
if command -v docker &> /dev/null; then
    print_status "Docker is available. You can test with Docker using:"
    echo "  docker-compose up --build"
    echo "  # or"
    echo "  docker build -f Dockerfile.freelancer -t answer-freelancer ."
    echo "  docker run -p 9080:80 answer-freelancer"
else
    print_warning "Docker not found. Skipping Docker tests."
fi

# Generate wire dependencies
print_status "Generating wire dependencies..."
if make generate; then
    print_success "Dependencies generated successfully!"
else
    print_error "Failed to generate dependencies!"
    exit 1
fi

# Build the application
print_status "Building the application..."
if make build; then
    print_success "Application built successfully!"
else
    print_error "Build failed!"
    exit 1
fi

# Check if the binary was created
if [ -f "answer" ]; then
    print_success "Binary 'answer' created successfully!"
    print_status "You can now run the application with: ./answer"
else
    print_error "Binary not found after build!"
    exit 1
fi

print_success "All tests completed successfully! 🎉"
echo ""
echo "Next steps:"
echo "1. Run the application: ./answer"
echo "2. Test API endpoints with curl or Postman"
echo "3. Test the UI components in the browser"
echo "4. Use Docker for containerized testing"
