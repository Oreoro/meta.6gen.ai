#!/bin/bash

# API Testing Examples for Freelancer Features
# This script provides curl examples to test your freelancer API endpoints

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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

# Configuration
BASE_URL="http://localhost:9080"
API_BASE="${BASE_URL}/answer/api/v1"

print_status "API Testing Examples for Freelancer Features"
echo "=================================================="
echo ""

print_status "Make sure your application is running on ${BASE_URL}"
print_status "You can start it with: ./answer"
echo ""

# Test 1: Create Freelancer Profile
print_status "Test 1: Create Freelancer Profile"
echo "----------------------------------------"
cat << 'EOF'
curl -X POST "${API_BASE}/freelancer/profile" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "is_available": true,
    "hourly_rate": 50.0,
    "currency": "USD",
    "skills": ["Go", "React", "Docker", "Kubernetes"],
    "experience": "5+ years of full-stack development experience",
    "portfolio": ["https://github.com/username/project1", "https://portfolio.example.com"],
    "availability": "Full-time",
    "preferred_projects": ["Web Development", "API Development", "DevOps"],
    "contact_email": "freelancer@example.com",
    "linkedin_profile": "https://linkedin.com/in/username",
    "github_profile": "https://github.com/username",
    "website": "https://freelancer.example.com",
    "bio": "Experienced full-stack developer with expertise in Go, React, and cloud technologies.",
    "languages": ["English", "Spanish"],
    "time_zone": "UTC-5",
    "response_time": "Within 24 hours"
  }'
EOF
echo ""

# Test 2: Update Freelancer Profile
print_status "Test 2: Update Freelancer Profile"
echo "---------------------------------------"
cat << 'EOF'
curl -X PUT "${API_BASE}/freelancer/profile" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "id": "profile_id_here",
    "is_available": false,
    "hourly_rate": 75.0,
    "currency": "EUR",
    "skills": ["Python", "Django", "PostgreSQL"],
    "experience": "Updated experience description",
    "portfolio": ["https://newproject.example.com"],
    "availability": "Part-time",
    "preferred_projects": ["Data Science", "Machine Learning"],
    "contact_email": "newemail@example.com",
    "linkedin_profile": "https://linkedin.com/in/newprofile",
    "github_profile": "https://github.com/newprofile",
    "website": "https://newwebsite.com",
    "bio": "Updated bio with new information",
    "languages": ["English", "French"],
    "time_zone": "UTC+1",
    "response_time": "Within 12 hours"
  }'
EOF
echo ""

# Test 3: Get Freelancer Profile
print_status "Test 3: Get Freelancer Profile"
echo "------------------------------------"
cat << 'EOF'
curl -X GET "${API_BASE}/freelancer/profile" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
EOF
echo ""

# Test 4: Get All Freelancer Profiles (for browsing)
print_status "Test 4: Get All Freelancer Profiles"
echo "------------------------------------------"
cat << 'EOF'
curl -X GET "${API_BASE}/freelancer/profiles?page=1&page_size=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
EOF
echo ""

# Test 5: Create Job Posting
print_status "Test 5: Create Job Posting"
echo "-------------------------------"
cat << 'EOF'
curl -X POST "${API_BASE}/freelancer/jobs" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Senior Go Developer Needed",
    "description": "We are looking for an experienced Go developer to join our team for a 6-month project.",
    "budget": 15000.0,
    "currency": "USD",
    "budget_type": "fixed",
    "skills": ["Go", "PostgreSQL", "Docker", "Kubernetes"],
    "experience_level": "senior",
    "duration": "6 months",
    "location": "remote",
    "contact_email": "hiring@company.com",
    "expires_at": "2024-12-31T23:59:59Z"
  }'
EOF
echo ""

# Test 6: Apply to Job
print_status "Test 6: Apply to Job"
echo "-----------------------"
cat << 'EOF'
curl -X POST "${API_BASE}/freelancer/jobs/JOB_ID/applications" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "cover_letter": "I am very interested in this position and believe my experience in Go development makes me a perfect fit.",
    "proposed_rate": 60.0,
    "currency": "USD",
    "message": "I can start immediately and am available for the full 6-month duration."
  }'
EOF
echo ""

# Test 7: Get Job Applications
print_status "Test 7: Get Job Applications"
echo "---------------------------------"
cat << 'EOF'
curl -X GET "${API_BASE}/freelancer/jobs/JOB_ID/applications" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
EOF
echo ""

print_warning "Note: Replace YOUR_JWT_TOKEN with an actual JWT token from your authentication system"
print_warning "Note: Replace JOB_ID with an actual job ID when testing job-related endpoints"
echo ""

print_status "To get a JWT token, you can:"
echo "1. Register a new user account"
echo "2. Login with existing credentials"
echo "3. Use the token from the login response"
echo ""

print_status "Example login request:"
cat << 'EOF'
curl -X POST "${API_BASE}/user/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "your_password"
  }'
EOF
echo ""

print_success "API testing examples ready! 🚀"
