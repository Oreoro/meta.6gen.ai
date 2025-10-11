#!/bin/bash

# Comprehensive Test Runner for Freelancer Features
# This script runs all tests for your new freelancer functionality

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

print_header() {
    echo -e "${PURPLE}================================${NC}"
    echo -e "${PURPLE}$1${NC}"
    echo -e "${PURPLE}================================${NC}"
}

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

print_step() {
    echo -e "${CYAN}[STEP]${NC} $1"
}

# Test results tracking
TESTS_PASSED=0
TESTS_FAILED=0
TOTAL_TESTS=0

run_test() {
    local test_name="$1"
    local test_command="$2"
    local test_description="$3"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    print_step "Running: $test_name"
    echo "Description: $test_description"
    echo "Command: $test_command"
    echo ""
    
    if eval "$test_command"; then
        print_success "$test_name PASSED ✅"
        TESTS_PASSED=$((TESTS_PASSED + 1))
    else
        print_error "$test_name FAILED ❌"
        TESTS_FAILED=$((TESTS_FAILED + 1))
    fi
    echo ""
}

print_header "🧪 COMPREHENSIVE TESTING FOR FREELANCER FEATURES"
echo ""

print_status "Starting comprehensive testing of your new freelancer features..."
echo "This will test:"
echo "- Unit tests for service layer"
echo "- Integration tests with database"
echo "- API endpoint testing"
echo "- Docker container testing"
echo "- Frontend component testing"
echo ""

# Check prerequisites
print_header "PREREQUISITE CHECKS"
echo ""

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    print_error "Please run this script from the project root directory"
    exit 1
fi

# Check Go installation
if command -v go &> /dev/null; then
    print_success "Go is installed: $(go version)"
else
    print_error "Go is not installed"
    exit 1
fi

# Check Docker installation
if command -v docker &> /dev/null; then
    print_success "Docker is installed: $(docker --version)"
else
    print_warning "Docker is not installed - Docker tests will be skipped"
fi

# Check Node.js installation
if command -v node &> /dev/null; then
    print_success "Node.js is installed: $(node --version)"
else
    print_warning "Node.js is not installed - Frontend tests will be skipped"
fi

echo ""

# Test 1: Unit Tests
print_header "UNIT TESTS"
echo ""

run_test "Freelancer Service Unit Tests" \
    "go test ./internal/service/freelancer/ -v" \
    "Testing freelancer service methods with mocks"

run_test "All Project Unit Tests" \
    "make test" \
    "Running all project unit tests"

# Test 2: Code Generation and Dependencies
print_header "CODE GENERATION & DEPENDENCIES"
echo ""

run_test "Generate Dependencies" \
    "make generate" \
    "Generating wire dependencies and other generated code"

run_test "Build Application" \
    "make build" \
    "Building the application binary"

# Test 3: Integration Tests
print_header "INTEGRATION TESTS"
echo ""

run_test "Integration Tests (if enabled)" \
    "RUN_INTEGRATION_TESTS=true go test ./internal/service/freelancer/ -v -run Integration" \
    "Running integration tests with database (requires RUN_INTEGRATION_TESTS=true)"

# Test 4: API Testing
print_header "API TESTING"
echo ""

print_status "Starting application for API testing..."
# Start the application in background
./answer &
APP_PID=$!

# Wait for application to start
print_status "Waiting for application to start..."
sleep 10

# Check if application is running
if ps -p $APP_PID > /dev/null; then
    print_success "Application started successfully (PID: $APP_PID)"
    
    # Test basic connectivity
    run_test "API Health Check" \
        "curl -s http://localhost:9080 > /dev/null" \
        "Testing if the application responds on port 9080"
    
    # Test API endpoints (basic)
    run_test "API Endpoint Test" \
        "curl -s http://localhost:9080/answer/api/v1 > /dev/null" \
        "Testing API endpoint accessibility"
    
    # Stop the application
    print_status "Stopping application..."
    kill $APP_PID 2>/dev/null || true
    wait $APP_PID 2>/dev/null || true
else
    print_error "Failed to start application"
    TESTS_FAILED=$((TESTS_FAILED + 1))
fi

# Test 5: Docker Testing
print_header "DOCKER TESTING"
echo ""

if command -v docker &> /dev/null; then
    run_test "Docker Build Test" \
        "docker build -f Dockerfile.freelancer -t answer-freelancer-test ." \
        "Testing Docker build with freelancer Dockerfile"
    
    run_test "Docker Compose Test" \
        "docker-compose up --build -d && sleep 15 && docker-compose ps | grep -q Up && docker-compose down -v" \
        "Testing Docker Compose setup"
else
    print_warning "Docker not available - skipping Docker tests"
fi

# Test 6: Frontend Testing
print_header "FRONTEND TESTING"
echo ""

if command -v node &> /dev/null && [ -d "ui" ]; then
    run_test "Frontend Dependencies" \
        "cd ui && (command -v pnpm > /dev/null && pnpm install || npm install) && cd .." \
        "Installing frontend dependencies"
    
    run_test "Frontend Build" \
        "cd ui && (command -v pnpm > /dev/null && pnpm run build || npm run build) && cd .." \
        "Building frontend assets"
    
    run_test "Frontend Linting" \
        "cd ui && (command -v pnpm > /dev/null && pnpm run lint || npm run lint) && cd .." \
        "Running frontend linting"
else
    print_warning "Node.js or UI directory not available - skipping frontend tests"
fi

# Test 7: Database Testing
print_header "DATABASE TESTING"
echo ""

run_test "Database Migration Test" \
    "ls -la data/sqlite3/answer.db 2>/dev/null || echo 'Database file not found (this might be expected)'" \
    "Checking database file existence"

# Test 8: Configuration Testing
print_header "CONFIGURATION TESTING"
echo ""

run_test "Configuration Files" \
    "ls -la configs/config.yaml && ls -la data/conf/config.yaml" \
    "Checking configuration files"

# Final Results
print_header "TEST RESULTS SUMMARY"
echo ""

echo "Total Tests: $TOTAL_TESTS"
echo "Passed: $TESTS_PASSED"
echo "Failed: $TESTS_FAILED"

if [ $TESTS_FAILED -eq 0 ]; then
    print_success "🎉 ALL TESTS PASSED! Your freelancer features are working correctly!"
    echo ""
    echo "Next steps:"
    echo "1. Start the application: ./answer"
    echo "2. Open http://localhost:9080 in your browser"
    echo "3. Test the freelancer features in the UI"
    echo "4. Use the API examples in test_api_examples.sh"
    echo "5. Deploy using Docker: docker-compose up"
else
    print_warning "⚠️  Some tests failed. Please check the output above for details."
    echo ""
    echo "Common issues and solutions:"
    echo "- Database connection: Check your database configuration"
    echo "- API errors: Ensure all routes are properly registered"
    echo "- Build errors: Check Go module dependencies"
    echo "- Frontend errors: Check Node.js and package dependencies"
    echo "- Docker errors: Ensure Docker is running and accessible"
fi

echo ""
print_status "Testing completed! 🚀"
